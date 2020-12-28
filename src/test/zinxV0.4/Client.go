package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

//模拟客户端
func main() {
	fmt.Println("client start...")
	time.Sleep(1 * time.Second)
	//连接远程server，获取一个conn
	dial, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Println("connect remote server error: ", err)
		return
	}
	//连接，写数据
	for {
		_, err := dial.Write([]byte("Hello Zinx V0.4..."))
		if err != nil {
			log.Println("write connection error", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := dial.Read(buf)
		if err != nil {
			log.Println("read buffer error", err)
			return
		}

		fmt.Printf("server call back : %s, cnt = %d\n", buf, cnt)
		//cpu 阻塞
		time.Sleep(1 * time.Second)
	}
}
