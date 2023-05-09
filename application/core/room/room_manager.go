package room

import (
	"Server/application/core"
	hasproto "Server/application/proto"
	"Server/utils"
	"Server/ziface"
	"Server/zlog"
	"errors"
	"fmt"
	"sync"
	"time"
)

/*
房间管理模块
*/
type RoomManager struct {
	Players map[int32]*core.Player //在房间内玩家的集合
	Waiting map[int32]*hasproto.SyncPlayerInfo
	pLock   sync.RWMutex //读写互斥锁
	wLock   sync.RWMutex
	logger  *zlog.Logger
}

var RoomMgrObj *RoomManager

func init() {
	RoomMgrObj = &RoomManager{
		Players: make(map[int32]*core.Player),
		Waiting: make(map[int32]*hasproto.SyncPlayerInfo),
		logger:  zlog.NewLogger("[RoomManager]"),
	}
}
func (rm *RoomManager) AddPlayer(player *core.Player) {
	rm.pLock.Lock()
	rm.Players[player.Pid] = player
	rm.pLock.Unlock()
}
func InitDefaultPlayer(PlayerInfo *hasproto.SyncPlayerInfo) *hasproto.BroadCast {
	DefaultBroadCast := &hasproto.BroadCast{
		Token:      PlayerInfo,
		Content:    utils.SpawnPlayer.DefaultContent,
		Info:       utils.SpawnPlayer.DefaultMovement,
		ActionData: utils.SpawnPlayer.DefaultAction,
	}
	return DefaultBroadCast
}
func (rm *RoomManager) RemovePlayerByPid(pid int32) {
	rm.pLock.Lock()
	delete(rm.Players, pid)
	rm.pLock.Unlock()
}
func (rm *RoomManager) KickPlayerByPid(pid int32) error {
	player := rm.GetPlayerByPid(pid)
	rm.RemovePlayerByPid(pid)
	players := rm.GetAllPlayers()
	if player != nil {
		for _, v := range players {
			err := v.KickPlayer(player.PlayerMessage)
			if err != nil {
				rm.logger.Error("Send Kick Player Info Error:" + err.Error())
				continue
			}
		}
	}
	return errors.New("err: No Player")
}

func (rm *RoomManager) SyncPlayerChat(has *hasproto.PlayerChat) error {
	players := rm.GetAllPlayers()
	for _, v := range players {
		if v != nil {
			err := v.SendChat(has)
			if err != nil {
				rm.logger.Error("Sync player chat error:" + err.Error())
				return err
			}
			return nil
		}
	}
	rm.logger.Error("No player founded")
	return errors.New("error:No players")
}

func (rm *RoomManager) GetPlayerByPid(pid int32) *core.Player {
	rm.pLock.Lock()
	defer rm.pLock.Unlock()
	return rm.Players[pid]
}

func (rm *RoomManager) GetAllPlayers() []*core.Player {
	rm.pLock.RLock()
	defer rm.pLock.RUnlock()
	players := make([]*core.Player, 0)
	for _, v := range rm.Players {
		players = append(players, v)
	}
	return players
}

func (rm *RoomManager) AddWaitingPlayer(pid int32, username string, token string) {
	playerInfo := &hasproto.SyncPlayerInfo{
		Pid: pid,
		Token: &hasproto.SyncPlayerToken{
			Token: token,
		},
		Name: username,
	}
	rm.wLock.Lock()
	rm.Waiting[pid] = playerInfo
	rm.wLock.Unlock()
	timer := time.NewTimer(5 * time.Second)
	go func(pid int32) {
		<-timer.C
		rm.wLock.Lock()
		delete(rm.Waiting, pid)
		rm.wLock.Unlock()
	}(pid)
}

func (rm *RoomManager) GetAllWaitingPlayer() []*hasproto.SyncPlayerInfo {
	rm.wLock.RLock()
	defer rm.wLock.RUnlock()
	players := make([]*hasproto.SyncPlayerInfo, 0)
	for _, v := range rm.Waiting {
		players = append(players, v)
	}
	return players
}
func (rm *RoomManager) FindWaitingPlayer(pid int32, token string) (*hasproto.SyncPlayerInfo, error) {
	rm.wLock.RLock()
	defer rm.wLock.RUnlock()
	var WaitingPlayer = rm.Waiting[pid]
	if WaitingPlayer != nil {
		if WaitingPlayer.Token.Token == token {
			return WaitingPlayer, nil
		}
		return nil, errors.New("token correct!")
	}
	return nil, errors.New("Can't find player!")
}

func (rm *RoomManager) MovePlayersFromWaitingToPlayers(PlayerInfo *hasproto.SyncPlayerInfo, conn ziface.IConnection) error {
	PlayerInfo, err := rm.FindWaitingPlayer(PlayerInfo.Pid, PlayerInfo.Token.Token)
	if err != nil {
		return err
	}
	DefaultBroadcast := InitDefaultPlayer(PlayerInfo)
	player := core.NewPlayer(conn, DefaultBroadcast)
	rm.AddPlayer(player)
	rm.logger.Info("Player:" + PlayerInfo.Name + ",UID:" + fmt.Sprint(PlayerInfo.Pid) + "  Login successed")
	return nil
}
