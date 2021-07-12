package game

import (
	"net"
)

//游戏状态
type RoomStatus uint8

const (
	Leisure RoomStatus = 1
	Game               = Leisure + 1
)

//房间
type Room struct {
	id         string                  //房间的唯一标识
	status     RoomStatus              //状态 0: 空闲中,1游戏中
	conn       *net.UDPConn            //udp服务器的唯一链接
	udpAddrMap map[string]*net.UDPAddr //加入房间的client
	join       chan *net.UDPAddr       //用户加入房间,  无缓冲，一个一个进入
	leave      chan *net.UDPAddr       //用户离开房间
	broadcast  chan []byte             //广播消息,无缓冲，要保证事件的顺序
}

//创建一个房间
func NewRoom(id string, conn *net.UDPConn) *Room {
	return &Room{
		id:         id,
		status:     Leisure,
		conn:       conn,
		udpAddrMap: make(map[string]*net.UDPAddr),
		join:       make(chan *net.UDPAddr),
		broadcast:  make(chan []byte),
	}
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
			g.udpAddrMap[join.String()] = join
		case leave := <-g.leave:
			//用户离开房间
			if _, ok := g.udpAddrMap[leave.String()]; ok {
				//删除这个用户
				delete(g.udpAddrMap, leave.String())
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

}
