package game

import (
	"net"
	"time"
)

const (
	//30s内判断用户是否活着
	AliveInterval = 30
)

//连接的玩家
type Player struct {
	udpAddr            *net.UDPAddr
	ready              bool  //是否已经准备
	lastAcceptPingTime int64 //上一次接收ping的时间戳，用来判断用户是否还活着，如果死亡了，则释放用户在房间的资源
}

func NewPlayer(udpAddr *net.UDPAddr) *Player {
	player := Player{udpAddr: udpAddr}
	player.SetLastAcceptPingTime(time.Now().Unix())
	return &player
}

//设置上一次接收ping的时间
func (p *Player) SetLastAcceptPingTime(lastAcceptPingTime int64) {
	p.lastAcceptPingTime = lastAcceptPingTime
}

//是否活着
func (p *Player) IsLive() bool {
	lastAcceptPingTime := time.Unix(p.lastAcceptPingTime, 0).Unix()
	nowTime := time.Now().Unix()

	if nowTime-lastAcceptPingTime > AliveInterval {
		//上一次存活时间大于30秒，则表示用户已经断块连接
		return false
	}
	return true
}
