package game

import (
	"errors"
	"net"
	"sync"
)

//房间管理者
type RoomManage struct {
	conn    *net.UDPConn     //udp服务器的唯一链接
	roomMap map[string]*Room //房间地图
	sync.RWMutex
}

func NewRoomManage(conn *net.UDPConn) *RoomManage {
	return &RoomManage{conn: conn, roomMap: make(map[string]*Room)}
}

//新建一个房间，成功返回true，失败返回false
func (m *RoomManage) AddRoom(roomId string) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.roomMap[roomId]; ok {
		return errors.New("房间已经存在")
	}

	m.roomMap[roomId] = NewRoom(roomId, m.conn)
	return nil
}

//删除一个房间
func (m *RoomManage) DeleteRoom(roomId string) {
	m.Lock()
	defer m.Unlock()

	if room, ok := m.roomMap[roomId]; ok {
		//房间资源释放
		room.Close()
		//删除map
		delete(m.roomMap, roomId)
	}

}

//获取房间
func (m *RoomManage) GetRoom(roomId string) (*Room, error) {
	m.RLock()
	defer m.RUnlock()

	if room, ok := m.roomMap[roomId]; ok {
		return room, nil
	}

	return nil, errors.New("房间找不到")
}
