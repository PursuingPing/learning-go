package znet

import (
	"errors"
	"io"
	"learning-go/src/zinx/utils"
	"learning-go/src/zinx/ziface"
	"log"
	"net"
)

/*
	连接模块
*/

type Connection struct {
	//当前connection隶属于那个server
	TcpServer ziface.IServer
	//当前连接的socket TCP
	Conn *net.TCPConn
	//连接id
	ConnID uint32
	//连接是否关闭
	isClosed bool
	//v0.2
	//处理当前业务的方法
	//handleAPI ziface.HandleFunc
	//告知当前连接已经退出的停止channel，由Reader告知Writer退出
	ExitChan chan bool
	//无缓冲的通道，用于读、写gorutine直接的消息通信
	msgChan chan []byte
	//消息的管理MsgID和对应得处理业务API关系
	MsgHandler ziface.IMsgHandler
}

//初始化连接模块的方法
func NewConnection(server ziface.IServer, conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandler) *Connection {
	c := &Connection{
		TcpServer:  server,
		Conn:       conn,
		ConnID:     connID,
		isClosed:   false,
		ExitChan:   make(chan bool, 1),
		msgChan:    make(chan []byte),
		MsgHandler: msgHandler,
	}
	//将connection加入到connectionManager中
	server.GetConnMgr().Add(c)
	return c
}

//连接的读业务方法
func (c *Connection) StartReader() {
	log.Println("[Reader Goroutine is running]...")
	defer log.Println("ConnId = ", c.ConnID, "【Reader is exit】, remote addr is ", c.GetRemoteAddr().String())
	defer c.Stop()

	for {
		//读取client的数据到buffer中
		//创建一个拆包对象
		dp := NewDataPack()

		//读取client的Msg的Head，二进制流
		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.GetTCPConnection(), headData)
		if err != nil {
			log.Println("[client close]read msg head error", err)
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
		//v0.7 go c.MsgHandler.DoMsgHandler(&req)
		//go func(request ziface.IMsgHandler) {
		//	c.Router.PreHandle(request)
		//	c.Router.Handle(request)
		//	c.Router.PostHandle(request)
		//}(&req)

		if utils.GlobalObject.WorkerPoolSize > 0 {
			//已经开启了workerPool
			//交给连接池处理
			c.MsgHandler.SendMsgToTaskQueue(&req)
		} else {
			go c.MsgHandler.DoMsgHandler(&req)
		}

	}
}

/*
	写消息的goroutine，专用于client回调处理
*/
func (c *Connection) StartWriter() {
	log.Println("[Writer goroutine] is running")
	defer log.Println("ConnId = ", c.ConnID, "[Writer Goroutine] is exit, remote addr is ", c.GetRemoteAddr().String())
	defer c.Stop()
	//不断地阻塞等待chan的消息，进行回写给client
	for {
		select {
		case data := <-c.msgChan:
			//有数据，回写给client（读写分离）
			if _, err := c.Conn.Write(data); err != nil {
				log.Println("Write data error", err)
				return
			}
		case <-c.ExitChan:
			//代表Reader已经退出，Writer也要退出
			log.Println("Reader ask Writer to quit ")
			return
		}
	}
}

func (c *Connection) Start() {
	log.Println("Conn Start()... ConnID = ", c.ConnID)
	//启动当前连接的读数据业务
	go c.StartReader()
	//启动写数据业务
	go c.StartWriter()

	//按照开发者传递进来的 创建连接之后的钩子函数
	c.TcpServer.CallOnConnStart(c)
}

func (c *Connection) Stop() {
	log.Println("Conn Stop()... ConnID = ", c.ConnID)

	if c.isClosed {
		return
	}
	c.isClosed = true

	//调用开发者注册的，销毁连接之前需要执行的业务钩子函数
	c.TcpServer.CallOnConnStop(c)
	//关闭socket连接
	err := c.Conn.Close()
	if err != nil {
		log.Println("Close socket connection error", err)
		return
	}
	//告知Writer关闭
	c.ExitChan <- true
	//将当前连接从manager中移除
	c.TcpServer.GetConnMgr().Remove(c)
	//关闭管道资源
	close(c.ExitChan)
	close(c.msgChan)
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
	//将数据发给Writer，让它发给客户端
	//if _, err := c.Conn.Write(binaryMsg); err != nil {
	//	log.Println("[Connection] conn Write msg id , error:", err)
	//	return errors.New("conn Write msg error ")
	//}
	c.msgChan <- binaryMsg
	return nil
}
