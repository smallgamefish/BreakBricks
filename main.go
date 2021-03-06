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
	//1Byte(字节) = 8bit(位)
	//1KB = 1024Byte(字节)
	//1MB = 1024KB
	//1GB = 1024MB
	//1TB = 1024GB
	MaxDataSize = 1024 * 1024
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

	roomManage := game.RoomManage
	roomManage.SetConn(socket)

	//监听用户链接
	for {
		//监听连接的用户
		data := make([]byte, MaxDataSize)
		n, remoteAddr, err := socket.ReadFromUDP(data)
		if err != nil {
			fmt.Println("读取数据失败!", err)
			continue
		}

		//事件解码
		log.Println("字节数是多少:", n)
		msg := new(protoc.ClientSendMsg)
		err = proto.Unmarshal(data[:n], msg)
		if err != nil {
			//解码失败
			log.Println("数据解码失败: ",err)
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
			log.Println("创建房间")
			//创建房间的事件
			err = roomManage.AddRoom(event.CreateRoomEvent.GetRoomId())

			response := new(protoc.ClientAcceptMsg)
			if err != nil {
				response.Code = protoc.ClientAcceptMsg_Error
				response.Error = err.Error()
			} else {
				response.Code = protoc.ClientAcceptMsg_Success
				response.Event = &protoc.ClientAcceptMsg_CreateRoomEvent{CreateRoomEvent: &protoc.CreateRoomEvent{
					RoomId: event.CreateRoomEvent.GetRoomId(),
				}}
			}

			responseData, _ := proto.Marshal(response)
			socket.WriteToUDP(responseData, remoteAddr)

		case *protoc.ClientSendMsg_JoinRoomEvent:
			//加入房间的事件
			err = roomManage.JoinRoom(event.JoinRoomEvent.GetRoomId(), remoteAddr)
			log.Println("加入房间:", err)

		case *protoc.ClientSendMsg_LeaveRoomEvent:
			//离开房间的事件
			err = roomManage.LeaveRoom(event.LeaveRoomEvent.GetRoomId(), remoteAddr)
			log.Println("离开房间:", err)

		case *protoc.ClientSendMsg_ReadyEvent:
			err = roomManage.ReadyRoom(event.ReadyEvent.GetRoomId(), remoteAddr, event.ReadyEvent.Ready)
			log.Println("用户在房间内准备:", err)
		case *protoc.ClientSendMsg_PingEvent:
			//客户端要定时发送ping给服务端，否则，服务器会认为这个用户已经断开了链接，会释放资源
			err = roomManage.UpdatePlayerLastAcceptPingTime(event.PingEvent.GetRoomId(), remoteAddr)
			log.Println("收到用户:",remoteAddr.String(),"的ping，确保用户链接还活着:", err)
		case *protoc.ClientSendMsg_FrameDataEvent:
			err = roomManage.BroadcastFrameData(event)
		}

	}
}
