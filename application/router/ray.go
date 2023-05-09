package router

import (
	"Server/application/core/gamemap"
	hasproto "Server/application/proto"
	"Server/ziface"
	"Server/zlog"
	"Server/znet"
)

type OnRayIntersect struct {
	znet.BaseRouter
	ray    *hasproto.RayIntersect
	Logger *zlog.Logger
}

func (ri *OnRayIntersect) PreHandle(req ziface.IRequest) {
	recvray, err := hasproto.UnmarshalRay(req)
	ri.ray = recvray
	if err != nil {
		ri.Logger.Error(err.Error())
	}
}
func (ri *OnRayIntersect) Handle(req ziface.IRequest) {
	ray := gamemap.Ray{Origin: gamemap.Corner{
		X: ri.ray.Origin.X,
		Y: ri.ray.Origin.Y,
		Z: ri.ray.Origin.Z,
	},
		Direction: gamemap.Corner{
			X: ri.ray.Direction.X,
			Y: ri.ray.Direction.Y,
			Z: ri.ray.Direction.Z,
		},
	}
	isRay := gamemap.Maputils.ActorRayIntersect(ri.ray.Actor, ray)
	if isRay {
		ri.Logger.Info("射中了" + ri.ray.Actor + "!")
	} else {
		ri.Logger.Info("没射中！" + ri.ray.Actor + "!")
	}
}
func (ri *OnRayIntersect) PostHandle(req ziface.IRequest) {

}
