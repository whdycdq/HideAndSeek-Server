package router

import (
	"Server/application/core/room"
	hasproto "Server/application/proto"
	"Server/ziface"
	"Server/znet"
	"fmt"
)

type OnPlayerLogin struct {
	znet.BaseRouter
	Info *hasproto.BroadCast
}

func (this *OnPlayerLogin) PreHandle(req ziface.IRequest) {
	info, err := hasproto.UnmarshalBroadCast(req)
	if err != nil {
		fmt.Println(err)
	}
	this.Info = info
}

func (this *OnPlayerLogin) Handle(req ziface.IRequest) {
	room.RoomMgrObj.MovePlayersFromWaitingToPlayers(this.Info.Token, req.GetConnection())
}

func (this *OnPlayerLogin) PostHandle(req ziface.IRequest) {

}
