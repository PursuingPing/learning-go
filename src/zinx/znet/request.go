package znet

import "learning-go/src/zinx/ziface"

/**
实现IRequest
*/
type Request struct {
	//已经与client建立好的连接
	conn ziface.IConnection
	//client 请求的数据
	data []byte
}

//得到当前连接
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

//得到请求数据
func (r *Request) GetData() []byte {
	return r.data
}
