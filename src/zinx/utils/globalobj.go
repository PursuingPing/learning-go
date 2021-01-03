package utils

import (
	"encoding/json"
	"io/ioutil"
	"learning-go/src/zinx/ziface"
)

/**
存储zinx的全局配置参数
*/

type GlobalObj struct {
	/*
		server
	*/
	TCPServer ziface.IServer
	Host      string
	TCPPort   int
	Name      string
	/*
		zinx
	*/
	Version          string
	MaxConn          int
	MaxPackageSize   uint32
	WorkerPoolSize   uint32 //当前业务工作Worker池的Goroutine数量
	MaxWorkerTaskLen uint32 //Zinx框架允许用户最多开辟多少个Worker（限制条件）
}

//定义一个全局的对外GlobalObj
var GlobalObject *GlobalObj

//加载自定义的参数
func (g *GlobalObj) Reload() {
	//读取文件,解析文件
	data, err := ioutil.ReadFile("src/test/zinxV0.9/conf/zinx.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

func init() {
	GlobalObject = &GlobalObj{
		Name:             "ZinxServerApp",
		Version:          "V0.5",
		TCPPort:          8999,
		Host:             "0.0.0.0",
		MaxConn:          1000,
		MaxPackageSize:   4096,
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024, //每个worker对应的消息队列的最大值
	}
	GlobalObject.Reload()
}
