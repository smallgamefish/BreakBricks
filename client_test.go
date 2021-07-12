package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/smallgamefish/BreakBricks/protoc/github.com/smallgamefish/BreakBricks/protoc"
	"log"
	"net"
	"testing"
)

func TestCreateRoomEvent(t *testing.T) {

	//拨号
	socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8100,
	})
	if err != nil {
		fmt.Println("连接失败!", err)
		return
	}

	//发送一个事件
	request := &protoc.ClientSendMsg{Event: &protoc.ClientSendMsg_CreateRoomEvent{
		CreateRoomEvent: &protoc.CreateRoomEvent{CreateRoomId: "1"},
	}}

	requestData, _ := proto.Marshal(request)

	_, err = socket.Write(requestData)
	if err != nil {
		log.Println("发送创建房间的事件失败", err)
	}

	//读取数据
	data := make([]byte, 1024*1024)
	n, _, err := socket.ReadFromUDP(data)
	if err != nil {
		log.Fatalln("读取数据失败:", err)
	}

	response := new(protoc.ClientAcceptMsg)
	err = proto.Unmarshal(data[:n], response)
	if err != nil {
		log.Fatalln("解码失败")
	}
	log.Println(response.Code.String())
	log.Println(response.Error)

}
