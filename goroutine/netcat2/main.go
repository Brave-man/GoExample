package main

import (
	"io"
	"log"
	"net"
	"os"
)

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// 第二个goroutine读取服务器的输入并输出
	go mustCopy(os.Stdout, conn)
	// 主goroutine从标准输入读取并发送到服务器
	mustCopy(conn, os.Stdin)
}
