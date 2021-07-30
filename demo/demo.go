package main

import (
	"flag"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/smallgamefish/BreakBricks/game"
	"github.com/smallgamefish/BreakBricks/protoc/github.com/smallgamefish/BreakBricks/protoc"
	"log"
	"net"
	"strings"
	"time"
)

const MaxDataSize = 1024 * 1024

func dial() *net.UDPConn {
	//拨号
	socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8100,
	})
	if err != nil {
		log.Fatalln("拨号失败:", err)
	}

	return socket
}

//创建房间
func createRoom(roomId string) {

	log.Println("创建房间的roomId:", roomId)

	socket := dial()

	//创建房间
	createRoomRequest := &protoc.ClientSendMsg{Event: &protoc.ClientSendMsg_CreateRoomEvent{
		CreateRoomEvent: &protoc.CreateRoomEvent{RoomId: roomId},
	}}

	createRoomData, _ := proto.Marshal(createRoomRequest)
	_, err := socket.Write(createRoomData)
	if err != nil {
		log.Fatalln("发送创建房间的事件失败", err)
	}

	//读取创建房间的结果
	data := make([]byte, 1024*1024)
	n, _, err := socket.ReadFromUDP(data)
	if err != nil {
		log.Fatalln("读取数据失败:", err)
	}

	createRoomResponse := new(protoc.ClientAcceptMsg)
	err = proto.Unmarshal(data[:n], createRoomResponse)
	if err != nil {
		log.Fatalln("解码失败")
	}

	if createRoomResponse.GetCode() != protoc.ClientAcceptMsg_Success {
		log.Fatalln("创建房间失败")
	}

	socket.Close()

	//加入房间
	joinRoom(roomId)
}

//加入房间
func joinRoom(roomId string) {
	log.Println("加入房间的roomId:", roomId)

	socket := dial()
	joinRoomRequest := &protoc.ClientSendMsg{Event: &protoc.ClientSendMsg_JoinRoomEvent{
		JoinRoomEvent: &protoc.JoinRoomEvent{RoomId: roomId},
	}}

	joinRoomData, _ := proto.Marshal(joinRoomRequest)
	_, err := socket.Write(joinRoomData)
	if err != nil {
		log.Fatalln("发送加入事件失败", err)
	}

	tick := time.Tick(3 * time.Second)

	//3s后发送准备事件
	go func() {
		<-tick

		readyRequest := &protoc.ClientSendMsg{Event: &protoc.ClientSendMsg_ReadyEvent{
			ReadyEvent: &protoc.ReadyEvent{RoomId: roomId, Ready: true},
		}}

		readyData, _ := proto.Marshal(readyRequest)
		_, err := socket.Write(readyData)
		if err != nil {
			log.Fatalln("发送准备事件失败:", err)
		}
		log.Println("发送准备事件完毕")

	}()

	//发送心跳检测，让服务端知道你还活着
	pingTick := time.Tick(game.AliveInterval / 2 * time.Second)
	go func() {
		for {
			select {
			case <-pingTick:
				readyRequest := &protoc.ClientSendMsg{Event: &protoc.ClientSendMsg_PingEvent{
					PingEvent: &protoc.PingEvent{RoomId: roomId},
				}}
				readyData, _ := proto.Marshal(readyRequest)
				socket.Write(readyData)
			}
		}
	}()

	//接受数据
	for {
		accept := make([]byte, MaxDataSize)
		n, _, err := socket.ReadFromUDP(accept)
		if err != nil {
			log.Fatalln("udp读取数据失败:", err)
		}

		msg := new(protoc.ClientAcceptMsg)
		err = proto.Unmarshal(accept[:n], msg)
		if err != nil {
			log.Fatalln("数据解码失败:", err)
		}

		switch event := msg.Event.(type) {
		case *protoc.ClientAcceptMsg_RefreshRoomPlayerEvent:
			//刷新房间的事件
			log.Println("房间用户刷新开始：")
			for _, player := range event.RefreshRoomPlayerEvent.GetPlayers() {
				log.Println("玩家地址udp唯一标识", player.GetUdpString())
				log.Println("玩家准备状态", player.GetReady())
			}
			log.Println("房间用户刷新完毕")
		case *protoc.ClientAcceptMsg_StartGameEvent:
			//监听游戏开始事件
			log.Println("开始游戏：", event.StartGameEvent.GetName())

			//触发发送数据帧
			frameDataSend(roomId, socket)

		case *protoc.ClientAcceptMsg_PingEvent:
			//心跳检测
			log.Println("接收到服务端推送过来的心跳检测：", event.PingEvent.RoomId)
		case *protoc.ClientAcceptMsg_FrameDataEvent:
			log.Println("收到数据帧", string(event.FrameDataEvent.GetFrameData()))
		}
	}
}

func frameDataSend(roomId string, socket *net.UDPConn) {

	tick := time.Tick(time.Second)
	go func() {
		i := 0
		for {
			select {
			case <-tick:

				readyRequest := &protoc.ClientSendMsg{Event: &protoc.ClientSendMsg_FrameDataEvent{
					FrameDataEvent: &protoc.FrameDataEvent{RoomId: roomId, FrameData: []byte(fmt.Sprintf("我是%s用户，现在发送第%d个数据帧", socket.LocalAddr().String(), i))},
				}}
				readyData, _ := proto.Marshal(readyRequest)
				socket.Write(readyData)
				i++
			}
		}
	}()
}

func main() {

	isJoinRoom := flag.Bool("isJoinRoom", false, "true加入房间，false创建房间，默认是false创建房间")
	roomId := flag.String("roomId", "", "房间id")

	flag.Parse()

	if strings.Compare("", *roomId) == 0 {
		log.Fatalln("房间号必填")
	}

	if *isJoinRoom {
		joinRoom(*roomId)
	} else {
		createRoom(*roomId)
	}

}
