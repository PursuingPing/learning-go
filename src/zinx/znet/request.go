package znet

import "learning-go/src/zinx/ziface"

/**
实现IRequest
*/
type Request struct {
	//已经与client建立好的连接
	conn ziface.IConnection
	//client 请求的数据
	msg ziface.IMessage
}

//得到当前连接
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

//得到请求数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

//得到请求的消息的id
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgId()
}
