package main

import (
	"Server/application/console"
	"Server/application/core/room"
	"Server/application/http"
	hasproto "Server/application/proto"
	"Server/application/router"
	"Server/utils"
	"Server/ziface"
	"Server/zlog"
	"Server/znet"
	"fmt"
)

/*
初始化玩家数据，使其处于conf设置里的状态，也就是基础状态
*/
func InitDefaultPlayerWithPidAndName(token string) *hasproto.BroadCast {
	tokenproto := &hasproto.SyncPlayerToken{
		Token: token,
	}
	PlayerInfo := &hasproto.SyncPlayerInfo{
		Token: tokenproto,
	}
	DefaultBroadCast := &hasproto.BroadCast{
		Token:      PlayerInfo,
		Content:    utils.SpawnPlayer.DefaultContent,
		Info:       utils.SpawnPlayer.DefaultMovement,
		ActionData: utils.SpawnPlayer.DefaultAction,
	}
	return DefaultBroadCast
}
func OnConnected(conn ziface.IConnection) {

}

func OnConnectionClosed(conn ziface.IConnection) {
	pid, err := conn.GetProperty("pid")
	if err != nil {
		fmt.Println("err:Cannot get pid")
		return
	}
	room.RoomMgrObj.KickPlayerByPid(pid.(int32))
}

func main() {
	go http.StartHTTPServer(8080)
	go console.App.Run()
	s := znet.NewServer()
	s.SetOnConnStart(OnConnected)
	s.SetOnConnStop(OnConnectionClosed)
	s.AddRouter(200, &router.OnBroadCast{Logger: zlog.NewLogger("[Broadcast]")})
	s.AddRouter(2, &router.OnSyncplayer{Logger: zlog.NewLogger("[SyncPlayer]")})
	s.AddRouter(202, &router.OnChat{Logger: zlog.NewLogger("[Chat]")})
	s.AddRouter(203, &router.OnRayIntersect{Logger: zlog.NewLogger("[RayIntersect]")})
	s.Serve()
}
