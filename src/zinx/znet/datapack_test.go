package znet

import (
	"io"
	"log"
	"net"
	"testing"
)

/**
封装、拆包模块测试用例
*/

func TestDataPack(t *testing.T) {
	/**
	模拟的服务器
	*/
	//创建socket
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Println("server listen error", err)
		return
	}
	go func() {
		//从client读取数据
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Println("server accept error", err)
				continue
			}
			go func(conn net.Conn) {
				//处理client端的请求
				//拆包过程
				dp := NewDataPack()
				for {
					//先从conn把包的head读处理
					headData := make([]byte, dp.GetHeadLen())
					_, err2 := io.ReadFull(conn, headData)
					if err2 != nil {
						log.Println("read head error", err2)
						break
					}
					message, err := dp.Unpack(headData)
					if err != nil {
						log.Println("server unpack error", err)
						return
					}
					//在从conn，根据head中的dataLen读取data的内容
					if message.GetMsgLen() > 0 {
						//msg是有数据的，需要进行第二次读取
						msg := message.(*Message)
						msg.Data = make([]byte, msg.GetMsgLen())
						//根据datalen的长度再次从io流中读取
						_, err := io.ReadFull(conn, msg.Data)
						if err != nil {
							log.Println("server unpack data error", err)
							return
						}
						//读取完毕
						log.Println("---> Receive MgsID: ", msg.Id, " dataLen= ", msg.DataLen, " data:", string(msg.Data))
					}
				}
			}(conn)
		}
	}()

	/**
	模拟client
	*/
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Println("client connection build error", err)
		return
	}
	//创建一个封包对象dp
	dp := NewDataPack()

	//模拟粘包过程，封装两个msg一起发送
	//第一个msg
	msg1 := &Message{
		Id:      1,
		DataLen: 4,
		Data:    []byte{'z', 'i', 'n', 'x'},
	}
	sendDataPack1, err := dp.Pack(msg1)
	if err != nil {
		log.Println("client pack massage1 error", err)
		return
	}
	//第二个msg
	msg2 := &Message{
		Id:      2,
		DataLen: 5,
		Data:    []byte{'H', 'e', 'l', 'l', 'o'},
	}
	sendDataPack2, err := dp.Pack(msg2)
	if err != nil {
		log.Println("client pack massage1 error", err)
		return
	}
	//粘包
	sendDataPack1 = append(sendDataPack1, sendDataPack2...)

	//一次性发送
	_, err2 := conn.Write(sendDataPack1)
	if err2 != nil {
		log.Println("send data error", err)
	}

	//阻塞，等待返回
	select {}
}
