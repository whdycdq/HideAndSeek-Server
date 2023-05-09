package router

import (
	"Server/application/core/room"
	hasproto "Server/application/proto"
	"Server/ziface"
	"Server/zlog"
	"Server/znet"
)

type OnChat struct {
	znet.BaseRouter
	chat   *hasproto.PlayerChat
	Logger *zlog.Logger
}

func (ct *OnChat) PreHandle(req ziface.IRequest) {
	recvchat, err := hasproto.UnmashalChat(req)
	ct.chat = recvchat
	if err != nil {
		ct.Logger.Error(err.Error())
	}
}
func (ct *OnChat) Handle(req ziface.IRequest) {
	if ct.chat.Type == 0 || ct.chat.Type == 1 {
		if len(ct.chat.Pid) != 0 {
			player := room.RoomMgrObj.GetPlayerByPid(ct.chat.Pid[1])
			if player != nil {
				err := player.SendChat(ct.chat)
				if err != nil {
					ct.Logger.Error(err.Error())
				}
			}
		}
	} else {
		err := room.RoomMgrObj.SyncPlayerChat(ct.chat)
		if err != nil {
			ct.Logger.Error(err.Error())
		}
	}

}

func (ct *OnChat) PostHandle(req ziface.IRequest) {
}
