package main

import (
	"fmt"
	"io"
	"learning-go/src/zinx/znet"
	"log"
	"net"
	"time"
)

//模拟客户端
func main() {
	fmt.Println("client1 start...")
	time.Sleep(1 * time.Second)
	//连接远程server，获取一个conn
	dial, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Println("connect remote server error: ", err)
		return
	}
	//连接，写数据
	for {
		//发送封装包的数据
		dp := znet.NewDataPack()
		binaryMsg, err := dp.Pack(znet.NewMessage(1, []byte("ZinxV0.8 client1 test Message")))
		if err != nil {
			log.Println("Pack error", err)
			return
		}
		if _, err := dial.Write(binaryMsg); err != nil {
			log.Println("write error", err)
			return
		}

		//服务器返回一个包
		//解析包
		//先读取head部分
		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(dial, binaryHead); err != nil {
			log.Println("client read msg head error", err)
			break
		}
		//在读取data
		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			log.Println("client unpack error", err)
			break
		}
		//根据dataLen再次读取 data，放入Message对象的data字段里
		var data []byte
		if msgHead.GetMsgLen() > 0 {
			data = make([]byte, msgHead.GetMsgLen())
			if _, err := io.ReadFull(dial, data); err != nil {
				log.Println("client read msg data error", err)
				continue
			}
		}
		msgHead.SetData(data)
		fmt.Println("--->Client1 Recv Server Msg : ID = ", msgHead.GetMsgId(), ", len = ", msgHead.GetMsgLen(), ", data=", string(msgHead.GetData()))
		//cpu 阻塞
		time.Sleep(1 * time.Second)
	}
}
