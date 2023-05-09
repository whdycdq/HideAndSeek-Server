package api

type UserInfo struct {
	PID int32 `bson:"pid"`
}

// 注册信息，用于用户验证
type UserRegInfo struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}

type Counters struct {
	CounterID string `bson:"_id"`
	Counter   int32  `bson:"counter"`
}
