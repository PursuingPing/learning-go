package ziface

/**
路由的抽象接口，路由里的数据都是IRequest
*/
type IRouter interface {
	//处理业务之前
	PreHandle(request IRequest)
	//处理业务的主方法
	Handle(request IRequest)
	//处理业务之后
	PostHandle(request IRequest)
}
