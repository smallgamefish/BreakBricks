# BreakBricks
打砖块游戏后台



# 开始游戏流程

1. 用户a创建房间 A;不会保存udp的socket地址
2. a用户加入 A房间，保存a用户的socket地址，会广播RefreshRoomPlayerEvent事件给房间内所有用户
3. b用户加入 A房间，保存b用户的socket地址，会广播RefreshRoomPlayerEvent事件给房间内所有用户
4. 用户在房间内，发送准备事件，服务器监听用户的准备状态，如果用户都准备了，广播游戏开始游戏事件，进入游戏
5. 进入游戏后，前端传什么事件，我就返回什么事件就行
6. udp是无链接状态，前端成功进入房间后或有游戏中，需要定时，每15s发送一次PingEvent时间，让服务器知道玩家的网络是正常的，否则服务器将玩家视为网络断开，会关闭、释放该玩家的资源；如果房间内的所有玩家都断开，则还会释放房间的资源



# 使用案例参考demo/demo.go



a玩家创建房间，并且加入房间，定时给服务器发送心跳检测
```html
# --roomId 房间的id
go run demo.go --roomId=1
```

b玩家加入房间，定时给服务器发送心跳检测
```html
# --roomId  房间的id
# --isJoinRoom=true 加入房间
 go run demo.go --roomId=1 --isJoinRoom=true
```