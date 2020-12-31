package main

import (
	"fmt"
	"learning-go/src/zinx/ziface"
	"learning-go/src/zinx/znet"
	"log"
)

type PingRouter struct {
	znet.BaseRouter
}

//Test PreHandle
func (this *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle...")
	connection := request.GetConnection()
	_, err := connection.GetTCPConnection().Write([]byte("Before ping...\n"))
	if err != nil {
		log.Println("call back before ping error", err)
	}
}

// Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")
	connection := request.GetConnection()
	_, err := connection.GetTCPConnection().Write([]byte("ping... ping...\n"))
	if err != nil {
		log.Println("call back ping ping error", err)
	}
}

// Test PostHandle
func (this *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle...")
	connection := request.GetConnection()
	_, err := connection.GetTCPConnection().Write([]byte("after ping...\n"))
	if err != nil {
		log.Println("call back after ping error", err)
	}
}

func main() {
	//创建一个server
	s := znet.NewServer("[zinx V0.4]")
	//添加一个路由
	s.AddRouter(0, &PingRouter{})
	//启动一个server
	s.Serve()
}
