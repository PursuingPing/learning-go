package znet

import (
	"learning-go/src/zinx/ziface"
	"log"
)

/**
消息处理模块的实现
*/

type MsgHandle struct {
	//存放每个msgID对应的处理Router
	Apis map[uint32]ziface.IRouter
}

//创建MsgHandler方法
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
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
