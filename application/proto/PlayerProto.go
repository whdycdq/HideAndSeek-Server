package hasproto

import (
	"Server/ziface"
	"errors"
	"google.golang.org/protobuf/proto"
)

func UnMarshalMovementMessage(msg ziface.IRequest) (*PlayerMovementInfo, error) {
	var newdata = &PlayerMovementInfo{}
	if newdata != nil {
		err := proto.Unmarshal(msg.GetData(), newdata)
		if err != nil {
			return nil, err
		}
		return newdata, nil
	}
	return nil, errors.New("Can not spawn MovementMessage")
}

func UnMarshalPlayerInfo(msg ziface.IRequest) (*SyncPlayerInfo, error) {
	var newdata = &SyncPlayerInfo{}
	if newdata != nil {
		err := proto.Unmarshal(msg.GetData(), newdata)
		if err != nil {
			return nil, err
		}
		return newdata, nil
	}
	return nil, errors.New("Can not spawn Playerinfo")
}
func UnmarshalBroadCast(msg ziface.IRequest) (*BroadCast, error) {
	var newdata = &BroadCast{}
	if newdata != nil {
		err := proto.Unmarshal(msg.GetData(), newdata)
		if err != nil {
			return nil, err
		}
		return newdata, nil
	}
	return nil, errors.New("Can not spawn broadcast")
}
func UnmashalChat(msg ziface.IRequest) (*PlayerChat, error) {
	var newdata = &PlayerChat{}
	if newdata != nil {
		err := proto.Unmarshal(msg.GetData(), newdata)
		if err != nil {
			return nil, err
		}
		return newdata, nil
	}
	return nil, errors.New("Can not spawn chat")
}

func UnmarshalRay(msg ziface.IRequest) (*RayIntersect, error) {
	var newdata = &RayIntersect{}
	if newdata != nil {
		err := proto.Unmarshal(msg.GetData(), newdata)
		if err != nil {
			return nil, err
		}
		return newdata, nil
	}
	return nil, errors.New("Can not spawn RayIntersect")
}
