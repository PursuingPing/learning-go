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

	err := request.GetConnection().SendMsg(201, []byte("Hello...Zinx...0.9"))
	if err != nil {
		log.Println(err)
	}
}

//创建连接之后执行的Hook
func DoConnectionBegin(conn ziface.IConnection) {
	log.Println("===> DoConnectionBegin is Called")
	err := conn.SendMsg(202, []byte("DoConnectionBegin...Zinx...0.9"))
	if err != nil {
		log.Println(err)
	}
}

//关闭连接之前执行的Hook
func DoConnectionLost(conn ziface.IConnection) {
	log.Println("===> DoConnectionLost is Called")
	//
	log.Println("ConnID = ", conn.GetConnID(), " is Lost")
}

func main() {
	//创建一个server
	s := znet.NewServer("[zinx V0.9]")
	//注册钩子方法
	//函数的名称其实就是一个地址
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)
	//添加一个路由
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})
	//启动一个server
	s.Serve()
}
