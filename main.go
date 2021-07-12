package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/smallgamefish/BreakBricks/game"
	"github.com/smallgamefish/BreakBricks/protoc/github.com/smallgamefish/BreakBricks/protoc"
	"log"
	"net"
)

const (
	//最大的消息长度，默认是512字节
	MaxDataSize = 512
)

func main() {

	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8100,
	})

	if err != nil {
		log.Fatalln("启动game监听失败", err)
	}

	log.Println("启动服务")

	roomManage := game.NewRoomManage(socket)

	//监听用户链接
	for {
		//监听连接的用户
		data := make([]byte, MaxDataSize)
		_, remoteAddr, err := socket.ReadFromUDP(data)
		if err != nil {
			fmt.Println("读取数据失败!", err)
			continue
		}

		msg := new(protoc.ClientSendMsg)
		err = proto.Unmarshal(data, msg)
		if err != nil {
			//解码失败
			response := new(protoc.ClientAcceptMsg)
			response.Code = protoc.ClientAcceptMsg_Error
			response.Error = err.Error()
			responseData, _ := proto.Marshal(response)
			socket.WriteToUDP(responseData, remoteAddr)
			continue
		}

		//事件分发
		switch event := msg.Event.(type) {

		case *protoc.ClientSendMsg_CreateRoomEvent:
			//创建房间的事件
			err = roomManage.AddRoom(event.CreateRoomEvent.GetCreateRoomId())
			response := new(protoc.ClientAcceptMsg)

			if err != nil {
				response.Code = protoc.ClientAcceptMsg_Error
				response.Error = err.Error()
			} else {
				response.Code = protoc.ClientAcceptMsg_Success
				response.Event = &protoc.ClientAcceptMsg_CreateRoomEvent{CreateRoomEvent: &protoc.CreateRoomEvent{
					CreateRoomId: event.CreateRoomEvent.GetCreateRoomId(),
				}}
			}

			responseData, _ := proto.Marshal(response)
			socket.WriteToUDP(responseData, remoteAddr)
			continue
		}

		////玩家加入房间
		//gameRoom.JoinChan() <- remoteAddr
		//
		//
		////广播消息
		//gameRoom.broadcast <- data
	}
}
