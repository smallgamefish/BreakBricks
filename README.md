# BreakBricks
打砖块游戏后台



#开始游戏流程

1. 用户a创建房间 A;不会保存udp的socket地址

2. a用户加入 A房间，保存a用户的socket地址，会广播JoinRoomEvent事件给房间内所有用户
3. b用户加入 A房间，保存b用户的socket地址，会广播JoinRoomEvent事件给房间内所有用户