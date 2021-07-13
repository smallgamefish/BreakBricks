package game

import (
	"github.com/golang/protobuf/proto"
	"github.com/smallgamefish/BreakBricks/protoc/github.com/smallgamefish/BreakBricks/protoc"
	"net"
)

//游戏状态
type RoomStatus uint8

const (
	//等待人满
	Prepare RoomStatus = iota
	//人满
	Full = Prepare + 1
	//游戏中
	Game = Full + 1
)

//房间
type Room struct {
	id           string                       //房间的唯一标识
	status       RoomStatus                   //状态 0: 等待人满,1人满,2游戏中
	maxUserTotal uint8                        //房间最大的人员
	conn         *net.UDPConn                 //udp服务器的唯一链接
	udpAddrMap   map[string]*net.UDPAddr      //加入房间的client
	closeRoom    chan bool                    //关闭房间，无缓冲通道
	join         chan *net.UDPAddr            //用户加入房间,  无缓冲，一个一个进入
	leave        chan *net.UDPAddr            //用户离开房间，无缓冲，一个一个离开
	broadcast    chan *protoc.ClientAcceptMsg //广播消息,有缓冲
}

//创建一个房间
func NewRoom(id string, conn *net.UDPConn) *Room {
	return &Room{
		id:           id,
		status:       Prepare,
		maxUserTotal: 2,
		conn:         conn,
		udpAddrMap:   make(map[string]*net.UDPAddr),
		closeRoom:    make(chan bool),
		join:         make(chan *net.UDPAddr),
		leave:        make(chan *net.UDPAddr),
		broadcast:    make(chan *protoc.ClientAcceptMsg, 20),
	}
}

//加入房间
func (g *Room) GetJoinChan() chan<- *net.UDPAddr {
	return g.join
}

func (g *Room) GetLeaveChan() chan<- *net.UDPAddr {
	return g.leave
}

func (g *Room) GetBroadcastChan() chan<- *protoc.ClientAcceptMsg {
	return g.broadcast
}

//获取存活的用户列表
func (g *Room) GetPlayers() []*protoc.Player {
	players := make([]*protoc.Player, len(g.udpAddrMap))

	i := 0
	for udpString, _ := range g.udpAddrMap {
		player := new(protoc.Player)
		player.UdpString = udpString
		players[i] = player
		i++
	}

	return players
}

//运行房间
func (g *Room) Run() {
	for {
		select {
		case join := <-g.join:
			//用户加入房间
			if g.status == Prepare {
				//等待人满的状态
				g.udpAddrMap[join.String()] = join
				if len(g.udpAddrMap) == int(g.maxUserTotal) {
					//加入的刚好满员了则切换状态
					g.status = Full
				}

				//服务端主动广播一个刷新房间用户的事件
				refreshRoomPlayerEventBroadcast := new(protoc.ClientAcceptMsg)
				refreshRoomPlayerEventBroadcast.Code = protoc.ClientAcceptMsg_Success
				refreshRoomPlayerEventBroadcast.Event = &protoc.ClientAcceptMsg_RefreshRoomPlayerEvent{RefreshRoomPlayerEvent: &protoc.RefreshRoomPlayerEvent{Players: g.GetPlayers()}}

				g.GetBroadcastChan() <- refreshRoomPlayerEventBroadcast
			}
		case leave := <-g.leave:
			//用户离开房间
			if _, ok := g.udpAddrMap[leave.String()]; ok {
				//删除这个用户
				delete(g.udpAddrMap, leave.String())

				if g.status == Full {
					//满员状态的切换成等待状态
					g.status = Prepare
				}

				if g.status == Prepare {
					//如果房间是等待状态的
					//服务端主动广播一个刷新房间用户的事件
					refreshRoomPlayerEventBroadcast := new(protoc.ClientAcceptMsg)
					refreshRoomPlayerEventBroadcast.Code = protoc.ClientAcceptMsg_Success
					refreshRoomPlayerEventBroadcast.Event = &protoc.ClientAcceptMsg_RefreshRoomPlayerEvent{RefreshRoomPlayerEvent: &protoc.RefreshRoomPlayerEvent{Players: g.GetPlayers()}}
					g.GetBroadcastChan() <- refreshRoomPlayerEventBroadcast
				}
			}

			if len(g.udpAddrMap) == 0 {
				//所有用户都断开了,房间失效，关闭房间,释放资源
				g.Close()
			}

		case data := <-g.broadcast:
			//广播消息
			g.sendMsg(data)
			//判断一下chan是否有缓冲信息，如果有，消费完它
			n := len(g.broadcast)
			for i := 0; i < n; i++ {
				g.sendMsg(<-g.broadcast)
			}
			
		case <-g.closeRoom:
			//关闭房间
			g.Close()
			return
		}
	}
}

//发送消息给房间内的所有用户
func (g *Room) sendMsg(event *protoc.ClientAcceptMsg) {
	responseData, _ := proto.Marshal(event)
	for _, udpAddr := range g.udpAddrMap {
		_, err := g.conn.WriteToUDP(responseData, udpAddr)
		if err != nil {
			//有用户链接断开了
			g.GetLeaveChan() <- udpAddr
		}
	}
}

//释放一些资源
func (g *Room) Close() {
	defer func() {
		recover()
	}()

	g.udpAddrMap = nil
	close(g.join)
	close(g.broadcast)
}
