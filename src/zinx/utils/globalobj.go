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
	Version        string
	MaxConn        int
	MaxPackageSize uint32
}

//定义一个全局的对外GlobalObj
var GlobalObject *GlobalObj

//加载自定义的参数
func (g *GlobalObj) Reload() {
	//读取文件,解析文件
	data, err := ioutil.ReadFile("src/test/zinxV0.6/conf/zinx.json")
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
		Name:           "ZinxServerApp",
		Version:        "V0.5",
		TCPPort:        8999,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}
	GlobalObject.Reload()
}
