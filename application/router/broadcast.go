package router

import (
	"Server/application/core/room"
	"Server/ziface"
	"Server/zlog"
	"Server/znet"
)

type OnBroadCast struct {
	znet.BaseRouter
	Logger *zlog.Logger
}

func (bc *OnBroadCast) PreHandle(req ziface.IRequest) {

}
func (bc *OnBroadCast) Handle(req ziface.IRequest) {
	//获取所有玩家
	players := room.RoomMgrObj.GetAllPlayers()
	//广播方法，向所有玩家广播其数据
	for _, v := range players {
		v.Conn.SendMsg(200, req.GetData())
	}
}
func (bc *OnBroadCast) PostHandle(req ziface.IRequest) {

}
