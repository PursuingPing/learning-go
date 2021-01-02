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
}
