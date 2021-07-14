package game

import (
	"errors"
	"fmt"
	"net"
	"sync"
)

const (
	MaxRoomNumber = 500
)

//房间管理者
var RoomManage *roomManage

func init() {
	if RoomManage == nil {
		RoomManage = &roomManage{conn: nil, roomMap: make(map[string]*Room)}
	}
}

//房间管理者
type roomManage struct {
	conn    *net.UDPConn     //udp服务器的唯一链接
	roomMap map[string]*Room //房间地图
	sync.RWMutex
}

//设置链接
func (m *roomManage) SetConn(conn *net.UDPConn) {
	m.conn = conn
}

//新建一个房间，成功返回true，失败返回false
func (m *roomManage) AddRoom(roomId string) error {
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
func (m *roomManage) JoinRoom(roomId string, player *net.UDPAddr) error {

	room, err := m.GetRoom(roomId)
	if err != nil {
		return err
	}

	//加入房间
	room.getJoinChan() <- NewPlayer(player)

	return nil
}

//用户离开房间
func (m *roomManage) LeaveRoom(roomId string, player *net.UDPAddr) error {

	room, err := m.GetRoom(roomId)
	if err != nil {
		return err
	}

	room.getLeaveChan() <- NewPlayer(player)

	return nil
}

//用户准备
func (m *roomManage) ReadyRoom(roomId string, player *net.UDPAddr, ready bool) error {
	room, err := m.GetRoom(roomId)
	if err != nil {
		return err
	}

	readyPlayer := NewPlayer(player)
	readyPlayer.ready = ready

	room.getReadyChan() <- readyPlayer

	return nil
}

//ping一下用户，确保用户还链接正常
func(m *roomManage) UpdatePlayerLastAcceptPingTime(roomId string, player *net.UDPAddr) error{
	room, err := m.GetRoom(roomId)
	if err != nil {
		return err
	}

	readyPlayer := NewPlayer(player)
	room.getPingActivePlayerChan() <- readyPlayer
	return nil
}

//删除一个房间
func (m *roomManage) deleteRoom(roomId string) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.roomMap[roomId]; ok {
		//删除map
		delete(m.roomMap, roomId)
	}
}

//获取房间
func (m *roomManage) GetRoom(roomId string) (*Room, error) {
	m.RLock()
	defer m.RUnlock()

	if room, ok := m.roomMap[roomId]; ok {
		return room, nil
	}

	return nil, errors.New("房间找不到")
}
