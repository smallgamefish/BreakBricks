package main

import (
	"flag"
	"github.com/golang/protobuf/proto"
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

	//30s后发送准备事件
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
		case *protoc.ClientAcceptMsg_PingEvent:
			//心跳检测
			log.Println("心跳检测：", event.PingEvent.Time)
		}
	}
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
