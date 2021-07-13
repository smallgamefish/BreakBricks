package game

import (
	"errors"
	"fmt"
	"github.com/smallgamefish/BreakBricks/protoc/github.com/smallgamefish/BreakBricks/protoc"
	"net"
	"sync"
)

const (
	MaxRoomNumber = 500
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

	if len(m.roomMap) == MaxRoomNumber {
		return errors.New(fmt.Sprintf("服务器目前最多只支持%d个房间", MaxRoomNumber))
	}

	newRoom := NewRoom(roomId, m.conn)
	m.roomMap[roomId] = newRoom
	//启动房间
	go newRoom.Run()
	return nil
}

//用户加入房间
func (m *RoomManage) JoinRoom(roomId string, user *net.UDPAddr) error {
	m.RLock()
	defer m.RUnlock()

	room, err := m.GetRoom(roomId)
	if err != nil {
		return err
	}

	if room.isJoin() {
		//加入房间
		room.GetJoinChan() <- user
		//广播一个加入房间的消息
		response := new(protoc.ClientAcceptMsg)
		response.Code = protoc.ClientAcceptMsg_Success
		response.Event = &protoc.ClientAcceptMsg_JoinRoomEvent{JoinRoomEvent: &protoc.JoinRoomEvent{RoomId: roomId}}
		room.GetBroadcastChan() <- response
	}

	return errors.New("无法加入房间")
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
