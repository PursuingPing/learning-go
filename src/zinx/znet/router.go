package znet

import "learning-go/src/zinx/ziface"

/**
实现IRouter时，先嵌入这个基类，根据各自需求对这个基类方法进行重写
*/
type BaseRouter struct{}

//使用模板方法模式
func (br *BaseRouter) PreHandle(request ziface.IRequest) {}

func (br *BaseRouter) Handle(request ziface.IRequest) {}

func (br *BaseRouter) PostHandle(request ziface.IRequest) {}
