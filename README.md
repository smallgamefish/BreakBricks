# BreakBricks
打砖块游戏后台



#开始游戏流程

1. 用户a创建房间 A;不会保存udp的socket地址

2. a用户加入 A房间，保存a用户的socket地址，会广播RefreshRoomPlayerEvent事件给房间内所有用户
3. b用户加入 A房间，保存b用户的socket地址，会广播RefreshRoomPlayerEvent事件给房间内所有用户
4. 用户在房间内，发送准备事件，服务器监听用户的准备状态，如果用户都准备了，广播游戏开始事件