package znet

import (
	"learning-go/src/zinx/utils"
	"learning-go/src/zinx/ziface"
	"log"
	"net"
)

/*
	连接模块
*/

type Connection struct {
	//当前连接的socket TCP
	Conn *net.TCPConn
	//连接id
	ConnID uint32
	//连接是否关闭
	isClosed bool
	//v0.2
	//处理当前业务的方法
	//handleAPI ziface.HandleFunc
	//告知当前连接已经退出的停止channel
	ExitChan chan bool
	//该连接处理的方法Router
	Router ziface.IRouter
}

//初始化连接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		ExitChan: make(chan bool, 1),
		Router:   router,
	}
	return c
}

//连接的读业务方法
func (c *Connection) StartReader() {
	log.Println("Reader Goroutine is running...")
	defer log.Println("ConnId = ", c.ConnID, "Reader is exit, remote addr is ", c.GetRemoteAddr().String())
	defer c.Stop()

	for {
		//读取client的数据到buffer中
		buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		//v0.2
		//cnt, err := c.Conn.Read(buf)
		_, err := c.Conn.Read(buf)
		if err != nil {
			log.Println("receive buf error", err)
			continue
		}
		////v0.2调用当前连接所绑定的处理api
		//if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
		//	log.Println("ConnID", c.ConnID, "handle is error", err)
		//	break
		//}
		////v0.3从路由中，找到注册绑定的Conn对应的router调用

		//得到当前conn的request
		req := Request{
			conn: c,
			data: buf,
		}
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}
}

func (c *Connection) Start() {
	log.Println("Conn Start()... ConnID = ", c.ConnID)
	//启动当前连接的读数据业务
	go c.StartReader()
	//TODO 启动写数据业务
}

func (c *Connection) Stop() {
	log.Println("Conn Stop()... ConnID = ", c.ConnID)

	if c.isClosed {
		return
	}
	c.isClosed = true
	//关闭socket连接
	err := c.Conn.Close()
	if err != nil {
		log.Println("Close socket connection error", err)
		return
	}
	close(c.ExitChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

//获取远程客户端TCP的状态，ip，端口
func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//发送数据
func (c *Connection) Send(data []byte) error {
	return nil
}
