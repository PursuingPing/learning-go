package ziface

/**
消息管理抽象层
*/

type IMsgHandler interface {
	//调度/执行对应的Router消息处理方法
	DoMsgHandler(request IRequest)
	//为消息添加具体处理逻辑
	AddRouter(msgID uint32, router IRouter)
	//启动WorkerPool
	StartWorkerPool()
	//将消息发给任务队列处理
	SendMsgToTaskQueue(request IRequest)
}
