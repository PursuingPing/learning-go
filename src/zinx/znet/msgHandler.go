package znet

import (
	"learning-go/src/zinx/utils"
	"learning-go/src/zinx/ziface"
	"log"
)

/**
消息处理模块的实现
*/

type MsgHandle struct {
	//存放每个msgID对应的处理Router
	Apis map[uint32]ziface.IRouter
	//负责worker去任务的消息队列
	TaskQueue []chan ziface.IRequest
	//业务Worker池中的worker数目
	WorkerPoolSize uint32
}

//创建MsgHandler方法
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis:           make(map[uint32]ziface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
	}
}

func (mh *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	//从request中找到msgID
	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		log.Println("api msgID = ", request.GetMsgID(), " is Not Found Need Register")
		return
	}
	//根据msgID得到Router//调用Router
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

func (mh *MsgHandle) AddRouter(msgID uint32, router ziface.IRouter) {
	//当前msg绑定的API处理方法是否存在
	if _, ok := mh.Apis[msgID]; ok {
		log.Println("repeat api, msgID = ", msgID)
		return
	}
	mh.Apis[msgID] = router
	log.Println("Add api MsgID = ", msgID, "successfully")
}

//启动一个Worker工作池,只能发生一次，一个zinx框架只有一个worker池
func (mh *MsgHandle) StartWorkerPool() {
	//根据workerPoolSize分布开启Worker，每个worker用一个goroutine承载
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		//启动一个worker
		//给当前worker对应得chan初始化
		mh.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		//启动当前的worker，阻塞地等待消息从chan中传递进来
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}

//一个Worker工作流程
func (mh *MsgHandle) StartOneWorker(workerID int, taskQueue chan ziface.IRequest) {
	log.Println("[StartOneWorker] WorkerID = ", workerID, " is start ...")

	//不断阻塞
	for {
		select {
		//如果有消息过来
		case request := <-taskQueue:
			mh.DoMsgHandler(request)
		}
	}
}

//将msg交给TaskQueue，由Worker处理 //单体，非分布式
func (mh *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {
	//负载均衡
	//根据connectionId分配，更好的是根据requestID
	workerID := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	log.Println("Add ConnID = ", request.GetConnection().GetConnID(), " request MsgID = ", request.GetMsgID(), " to WorkerID = ", workerID)
	//发送
	mh.TaskQueue[workerID] <- request
}
