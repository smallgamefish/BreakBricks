package game

import (
	"net"
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
	id           string                  //房间的唯一标识
	status       RoomStatus              //状态 0: 空闲中,1准备等待人满,2满员,3游戏中
	maxUserTotal uint8                   //房间最大的人员
	conn         *net.UDPConn            //udp服务器的唯一链接
	udpAddrMap   map[string]*net.UDPAddr //加入房间的client
	join         chan *net.UDPAddr       //用户加入房间,  无缓冲，一个一个进入
	leave        chan *net.UDPAddr       //用户离开房间
	broadcast    chan []byte             //广播消息,无缓冲，要保证事件的顺序
}

//创建一个房间
func NewRoom(id string, conn *net.UDPConn) *Room {
	return &Room{
		id:           id,
		status:       Leisure,
		maxUserTotal: 2,
		conn:         conn,
		udpAddrMap:   make(map[string]*net.UDPAddr),
		join:         make(chan *net.UDPAddr),
		broadcast:    make(chan []byte),
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
func (g *Room) JoinChan() chan<- *net.UDPAddr {
	return g.join
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
			for _, udpAddr := range g.udpAddrMap {
				n, err := g.conn.WriteToUDP(data, udpAddr)
				if err != nil || n == 0 {
					continue
				}
			}
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
