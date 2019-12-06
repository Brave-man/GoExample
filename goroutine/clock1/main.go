package main

import (
	"io"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("2006-01-02 15:04:05\n"))
		if err != nil {
			return // 例如连接断开
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000") // 监听tcp 8000端口
	if err != nil {
		log.Fatal(err)
	}

	for {
		// Accept方法被阻塞，直到有请求进来
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// 处理一个完整的请求
		handleConn(conn)
	}
}
