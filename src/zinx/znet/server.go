package znet

import (
	"fmt"
	"learning-go/src/zinx/utils"
	"learning-go/src/zinx/ziface"
	"log"
	"net"
)

//IServer的接口实现，定义一个Server服务器模块
type Server struct {
	//名称
	Name string
	//ip版本
	IPVersion string
	//监听的ip
	Ip string
	//端口
	Port int
	//当前server的消息管理模块
	MsgHandler ziface.IMsgHandler
}

//v0.2
//func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
//	//回显的业务
//	log.Println("[Conn Handle] CallBackClient ...")
//	fmt.Printf("receive from client buf :%s, cnt %d\n", data, cnt)
//	if _, err := conn.Write(data[:cnt]); err != nil {
//		log.Println("write back buffer error", err)
//		return errors.New("CallBackToClient error")
//	}
//	return nil
//}

func (s *Server) Start() {
	log.Printf("[Zinx] Server Name:%s, listenning at IP: %s,Port:%d is starting\n",
		utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TCPPort)
	log.Printf("[Zinx] Version %s,MaxConnection: %d, MaxPackageSize: %d\n",
		utils.GlobalObject.Version, utils.GlobalObject.MaxConn, utils.GlobalObject.MaxPackageSize)
	//log.Printf("[Start] Server Listenner at IP: %s, Port: %d, is starting\n", s.Ip, s.Port)
	go func() {
		//获取一个TCP的addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.Ip, s.Port))
		if err != nil {
			log.Println("resolve tcp addr error :", err)
			return
		}

		//监听服务器的地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			log.Println("listen ", s.Ip, " error: ", err)
			return
		}
		log.Println("start Zinx server successfully", s.Name, "listening...")
		//阻塞地等待client连接，处理业务
		var cid uint32
		cid = 0
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				log.Println("Accept error", err)
				continue
			}
			//tcp连接成功
			//将处理业务的方法与connection绑定
			dealConn := NewConnection(conn, cid, s.MsgHandler)
			cid++
			//启动当前业务处理
			go dealConn.Start()
			//v0.1
			//go func() {
			//	for {
			//		buf := make([]byte, 512)
			//		cnt, err := conn.Read(buf)
			//		if err != nil {
			//			log.Println("receive buffer error", err)
			//			continue
			//		}
			//		fmt.Printf("receive from client buf :%s, cnt %d\n", buf, cnt)
			//		if _, err := conn.Write(buf[:cnt]); err != nil {
			//			log.Println("write back buffer error", err)
			//			continue
			//		}
			//	}
			//}()
		}
	}()
}

func (s *Server) Stop() {
	//TODO 将服务器资源、状态或一些连接信息停止与回收
}

func (s *Server) Serve() {
	//启动server的服务功能
	s.Start()
	//TODO 启动服务器之后的业务
	//阻塞状态
	select {}
}

//添加一个路由
func (s *Server) AddRouter(msgID uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgID, router)
	log.Println("Add Router successfully!!")
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:       name,
		IPVersion:  "tcp4",
		Ip:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TCPPort,
		MsgHandler: NewMsgHandle(),
	}
	return s
}
