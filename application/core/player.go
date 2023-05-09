package core

import (
	"Server/application/proto"
	"Server/ziface"
	"errors"
	"fmt"
	"google.golang.org/protobuf/proto"
	"sync"
)

// 在server上生成的玩家对象
type Player struct {
	Pid           int32               //玩家的id
	Conn          ziface.IConnection  //当前玩家的连接
	PlayerMessage *hasproto.BroadCast //玩家在游戏中所有可能同步的信息
}

var PidGen int32 = 100001
var IdLock sync.Mutex

func NewPlayer(conn ziface.IConnection, DefaultSpawn *hasproto.BroadCast) *Player {
	//TODO 生成一个PID，后期要从数据库获取id
	p := &Player{
		//暂时不初始化
		Pid:           DefaultSpawn.Token.Pid,
		Conn:          conn,
		PlayerMessage: DefaultSpawn,
	}
	return p
}

/*
发送消息给客户端，
主要是将pb的protobuf数据序列化之后发送
*/
func (p *Player) SendMessage(msgID uint32, data *hasproto.BroadCast) error {
	msg, err := proto.Marshal(data)
	fmt.Printf("长度为%d\n", len(msg))
	//	fmt.Printf("Token:%s\nUsername:%s\nPid:%d\n", data.Token.Token.Token, data.Token.Name, data.Token.Pid)
	if err != nil {
		return err
	}
	if p.Conn == nil {
		return errors.New("Can not find player connection!")
	}
	if err := p.Conn.SendMsg(msgID, msg); err != nil {
		return err
	}
	return nil
}
func (p *Player) SendToken(token []byte) error {
	err := p.Conn.SendMsg(1, token)
	if err != nil {
		return err
	}
	return nil
}

func (p *Player) Sync(req ziface.IRequest) error {
	err := p.Conn.SendMsg(200, req.GetData())
	if err != nil {
		return err
	}
	return nil
}

func (p *Player) SyncPlayerInfo(req ziface.IRequest) error {
	err := p.Conn.SendMsg(2, req.GetData())
	if err != nil {
		return err
	}
	return nil
}

func (p *Player) KickPlayer(has *hasproto.BroadCast) error {
	data, err := proto.Marshal(has)
	if err != nil {
		return err
	}
	err = p.Conn.SendMsg(201, data)
	if err != nil {
		return err
	}
	return nil
}

func (p *Player) SendChat(has *hasproto.PlayerChat) error {
	data, err := proto.Marshal(has)
	if err != nil {
		return err
	}
	err = p.Conn.SendMsg(202, data)
	if err != nil {
		return err
	}
	return nil
}
