package ziface

//定义一个服务器接口
type IServer interface {
	//启动
	Start()
	//停止
	Stop()
	//运行
	Serve()
	//集成router模块,给当前服务注册一个路由方法，供client的连接使用
	AddRouter(msgID uint32, router IRouter)

	GetConnMgr() IConnManager
	//注册连接开启之后的钩子函数
	SetOnConnStart(func(connection IConnection))
	//注册连接关闭的钩子函数
	SetOnConnStop(func(connection IConnection))
	//调用连接开启的钩子函数
	CallOnConnStart(connection IConnection)
	//调用连接关闭之前的钩子函数
	CallOnConnStop(connection IConnection)
}
