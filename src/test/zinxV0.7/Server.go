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

// Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")
	//先读取client的数据，在回写ping...ping...ping
	log.Println("recv from client: msgID= ", request.GetMsgID(), ", data= ", string(request.GetData()))

	err := request.GetConnection().SendMsg(200, []byte("ping...ping...ping"))
	if err != nil {
		log.Println(err)
	}
}

//1,hello router
type HelloZinxRouter struct {
	znet.BaseRouter
}

func (this *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")
	//先读取client的数据，在回写ping...ping...ping
	log.Println("recv from client: msgID= ", request.GetMsgID(), ", data= ", string(request.GetData()))

	err := request.GetConnection().SendMsg(201, []byte("Hello...Zinx...0.7"))
	if err != nil {
		log.Println(err)
	}
}

func main() {
	//创建一个server
	s := znet.NewServer("[zinx V0.7]")
	//添加一个路由
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})
	//启动一个server
	s.Serve()
}
