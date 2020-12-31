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

	err := request.GetConnection().SendMsg(1, []byte("ping...ping...ping"))
	if err != nil {
		log.Println(err)
	}
}

func main() {
	//创建一个server
	s := znet.NewServer("[zinx V0.5]")
	//添加一个路由
	s.AddRouter(0, &PingRouter{})
	//启动一个server
	s.Serve()
}
