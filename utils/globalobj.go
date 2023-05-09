package utils

import (
	hasproto "Server/application/proto"
	"Server/ziface"
	"encoding/json"
	"os"
)

/*
存储有关初始人物状态的结构体
*/
type DefaultSpawnPlayer struct {
	DefaultMovement    *hasproto.PlayerMovementInfo //默认出生位置和出生角度等
	DefaultContent     int32                        //默认出生状态
	DefaultAction      int32                        //默认出生动作
	PlayerConfFilePath string
}

/*
存储一切有关Zinx框架的全局参数，供其他模块使用
一些参数也可以通过 用户根据 zinx.json来配置
*/
type GlobalObj struct {
	/*
		Server
	*/
	TcpServer ziface.IServer //当前Zinx的全局Server对象
	Host      string         //当前服务器主机IP
	TcpPort   int            //当前服务器主机监听端口号
	Name      string         //当前服务器名称
	SecretKey string
	/*
		Zinx
	*/
	Version          string //当前Zinx版本号
	MaxPacketSize    uint32 //都需数据包的最大值
	MaxConn          int    //当前服务器主机允许的最大链接个数
	WorkerPoolSize   uint32 //业务工作Worker池的数量
	MaxWorkerTaskLen uint32 //业务工作Worker对应负责的任务队列最大任务存储数量
	MaxMsgChanLen    uint32 //SendBuffMsg发送消息的缓冲最大长度

	/*
		config file path
	*/
	ConfFilePath string
}

/*
定义一个全局的对象
*/
var GlobalObject *GlobalObj
var SpawnPlayer *DefaultSpawnPlayer

// 判断一个文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 读取用户的配置文件
func (g *GlobalObj) Reload() {

	if confFileExists, _ := PathExists(g.ConfFilePath); confFileExists != true {
		//fmt.Println("Config File ", g.ConfFilePath , " is not exist!!")
		return
	}

	data, err := os.ReadFile(g.ConfFilePath)
	if err != nil {
		panic(err)
	}
	//将json数据解析到struct中
	//fmt.Printf("json :%s\n", data)
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}
func (p *DefaultSpawnPlayer) PlayerReload() {
	if confFileExists, _ := PathExists(p.PlayerConfFilePath); confFileExists != true {
		//fmt.Println("Config File ", g.ConfFilePath , " is not exist!!")
		return
	}
	data, err := os.ReadFile(p.PlayerConfFilePath)
	if err != nil {
		panic(err)
	}
	temp := struct {
		LocationX float32
		LocationY float32
		LocationZ float32
		Content   int32
		Action    int32
	}{}
	err = json.Unmarshal(data, &temp)
	if err != nil {
		panic(err)
	}
	/*
		设置protobuf对象，以便直接传出
	*/
	DefaultLocation := hasproto.Location{
		X: temp.LocationX,
		Y: temp.LocationY,
		Z: temp.LocationZ,
	}
	DefaultRotation := hasproto.Rotation{
		Pitch: 0.0,
		Yaw:   0.0,
		Roll:  0.0,
	}
	DefaultScale := hasproto.Scale{
		X: 1.0,
		Y: 1.0,
		Z: 1.0,
	}
	DefaultTran := hasproto.Transform{
		L: &DefaultLocation,
		R: &DefaultRotation,
		S: &DefaultScale,
	}
	DefaultCRotator := hasproto.ControllerRotation{
		Pitch: 0.0,
		Yaw:   0.0,
		Roll:  0.0,
	}
	Axis := hasproto.Axis{
		MoveForward: 0.0,
		MoveRight:   0.0,
	}
	DefaultInfo := hasproto.PlayerMovementInfo{
		Trans:     &DefaultTran,
		CRotation: &DefaultCRotator,
		InputAxis: &Axis,
	}
	SpawnPlayer.DefaultMovement = &DefaultInfo
	SpawnPlayer.DefaultContent = temp.Content
	SpawnPlayer.DefaultAction = temp.Action
}

/*
提供init方法，默认加载
*/
func init() {
	//初始化GlobalObject变量，设置一些默认值
	GlobalObject = &GlobalObj{
		Name:             "ZinxServerApp",
		Version:          "V0.4",
		TcpPort:          7777,
		Host:             "0.0.0.0",
		MaxConn:          12000,
		MaxPacketSize:    4096,
		ConfFilePath:     "application/conf/server.json",
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024,
		MaxMsgChanLen:    1024,
	}

	SpawnPlayer = &DefaultSpawnPlayer{
		DefaultMovement:    &hasproto.PlayerMovementInfo{},
		DefaultContent:     1,
		DefaultAction:      1,
		PlayerConfFilePath: "application/conf/player.json",
	}

	//从配置文件中加载一些用户配置的参数
	GlobalObject.Reload()
	SpawnPlayer.PlayerReload()
}
