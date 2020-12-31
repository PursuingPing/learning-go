package znet

import (
	"errors"
	"io"
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
	//消息的管理MsgID和对应得处理业务API关系
	MsgHandler ziface.IMsgHandler
}

//初始化连接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandler) *Connection {
	c := &Connection{
		Conn:       conn,
		ConnID:     connID,
		isClosed:   false,
		ExitChan:   make(chan bool, 1),
		MsgHandler: msgHandler,
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
		//创建一个拆包对象
		dp := NewDataPack()

		//读取client的Msg的Head，二进制流
		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.GetTCPConnection(), headData)
		if err != nil {
			log.Println("read msg head error", err)
			break
		}
		//拆包，得到msgId与dataLen，放入Message对象中
		msg, err := dp.Unpack(headData)
		if err != nil {
			log.Println("unpack error", err)
			break
		}
		//根据dataLen再次读取 data，放入Message对象的data字段里
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.Conn, data); err != nil {
				log.Println("read msg data error", err)
				break
			}
		}
		msg.SetData(data)

		//buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		////v0.2
		////cnt, err := c.Conn.Read(buf)
		//_, err := c.Conn.Read(buf)
		//if err != nil {
		//	log.Println("receive buf error", err)
		//	continue
		//}
		////v0.2调用当前连接所绑定的处理api
		//if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
		//	log.Println("ConnID", c.ConnID, "handle is error", err)
		//	break
		//}
		////v0.3从路由中，找到注册绑定的Conn对应的router调用

		//得到当前conn的request
		req := Request{
			conn: c,
			msg:  msg,
		}
		go c.MsgHandler.DoMsgHandler(&req)
		//go func(request ziface.IMsgHandler) {
		//	c.Router.PreHandle(request)
		//	c.Router.Handle(request)
		//	c.Router.PostHandle(request)
		//}(&req)
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

//发送数据，包装数据成包
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	//如果当前c关闭了
	if c.isClosed == true {
		return errors.New("Connection is Closed when sending msg ")
	}
	//封包
	dp := NewDataPack()
	binaryMsg, err := dp.Pack(NewMessage(msgId, data))
	if err != nil {
		log.Println("pack error msg id = , error", msgId, err)
		return errors.New("Pack error msg ")
	}

	//发送
	if _, err := c.Conn.Write(binaryMsg); err != nil {
		log.Println("[Connection] conn Write msg id , error:", err)
		return errors.New("conn Write msg error ")
	}
	return nil
}
