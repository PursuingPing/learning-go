package main

import "learning-go/src/zinx/znet"

func main() {
	//创建一个server
	s := znet.NewServer("[zinx V0.1]")
	//启动一个server
	s.Serve()
}
