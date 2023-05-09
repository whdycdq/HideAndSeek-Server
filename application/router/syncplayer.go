package router

import (
	"Server/application/core/room"
	hasproto "Server/application/proto"
	"Server/ziface"
	"Server/zlog"
	"Server/znet"
	"google.golang.org/protobuf/proto"
)

type OnSyncplayer struct {
	znet.BaseRouter
	Info   *hasproto.BroadCast
	Logger *zlog.Logger
}

func (sy *OnSyncplayer) PreHandle(req ziface.IRequest) {
	info, err := hasproto.UnmarshalBroadCast(req)
	if err != nil {
		sy.Logger.Info(err.Error())
	}
	sy.Info = info
	req.GetConnection().SetProperty("pid", sy.Info.Token.Pid)
}
func (sy *OnSyncplayer) Handle(req ziface.IRequest) {
	err := room.RoomMgrObj.MovePlayersFromWaitingToPlayers(sy.Info.Token, req.GetConnection())
	if err != nil {
		return
	}
	player := room.RoomMgrObj.GetPlayerByPid(sy.Info.Token.Pid)
	players := room.RoomMgrObj.GetAllPlayers()
	for _, v := range players {
		temp, err := proto.Marshal(v.PlayerMessage)
		if err != nil {
			sy.Logger.Error(err.Error())
			break
		}
		err = player.Conn.SendMsg(2, temp)
		if v != player {
			v.SendMessage(2, player.PlayerMessage)
		}
		if err != nil {
			sy.Logger.Error(err.Error())
		}
	}
}
func (sy *OnSyncplayer) PostHandle(req ziface.IRequest) {

}
