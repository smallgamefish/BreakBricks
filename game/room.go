package game

import (
	"github.com/smallgamefish/BreakBricks/protoc/github.com/smallgamefish/BreakBricks/protoc"
	"net"
	"github.com/golang/protobuf/proto"
)

//游戏状态
type RoomStatus uint8

const (
	//空闲中
	Leisure RoomStatus = iota
	//准备中：等待人满
	Prepare = Leisure + 1
	//满员
	Full = Prepare + 1
	Game = Full + 1
)

//房间
type Room struct {
	id           string                       //房间的唯一标识
	status       RoomStatus                   //状态 0: 空闲中,1准备等待人满,2满员,3游戏中
	maxUserTotal uint8                        //房间最大的人员
	conn         *net.UDPConn                 //udp服务器的唯一链接
	udpAddrMap   map[string]*net.UDPAddr      //加入房间的client
	closeRoom    chan bool                    //关闭房间，无缓冲通道
	join         chan *net.UDPAddr            //用户加入房间,  无缓冲，一个一个进入
	leave        chan *net.UDPAddr            //用户离开房间
	broadcast    chan *protoc.ClientAcceptMsg //广播消息,无缓冲，要保证事件的顺序
}

//创建一个房间
func NewRoom(id string, conn *net.UDPConn) *Room {
	return &Room{
		id:           id,
		status:       Leisure,
		maxUserTotal: 2,
		conn:         conn,
		udpAddrMap:   make(map[string]*net.UDPAddr),
		closeRoom:    make(chan bool),
		join:         make(chan *net.UDPAddr),
		broadcast:    make(chan *protoc.ClientAcceptMsg),
	}
}

//判断是否可以加入房间
func (g *Room) isJoin() bool {
	if g.status == Prepare {
		return true
	}
	return false
}

//加入房间
func (g *Room) GetJoinChan() chan<- *net.UDPAddr {
	return g.join
}

func (g *Room) GetBroadcastChan() chan<- *protoc.ClientAcceptMsg {
	return g.broadcast
}

//运行房间
func (g *Room) Run() {
	for {
		select {
		case join := <-g.join:
			//用户加入房间
			if g.status == Prepare {
				//准备中房间可加用户
				g.udpAddrMap[join.String()] = join

				if len(g.udpAddrMap) == int(g.maxUserTotal) {
					//加入的刚好满员了则切换状态
					g.status = Full
				}
			}
		case leave := <-g.leave:
			//用户离开房间
			if _, ok := g.udpAddrMap[leave.String()]; ok {
				//删除这个用户
				delete(g.udpAddrMap, leave.String())
				//房间是满员状态的，切换到准备状态
				if g.status == Full {
					g.status = Prepare
				}
			}
		case data := <-g.broadcast:
			//广播消息
			responseData, _ := proto.Marshal(data)
			for _, udpAddr := range g.udpAddrMap {
				n, err := g.conn.WriteToUDP(responseData, udpAddr)
				if err != nil || n == 0 {
					continue
				}
			}

		case <-g.closeRoom:
			//关闭房间
			g.Close()
			return
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
	close(g.leave)
	close(g.broadcast)
}
