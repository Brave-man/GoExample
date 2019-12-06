package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string // 对外发送消息的通道

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // 所有接受的客户消息
)

// broadcaster 广播器
func broadcaster() {
	// 使用局部变量clients来记录当前连接的客户集合
	// 每个客户唯一被记录的信息是其对外发送消息通道的ID
	clients := make(map[client]bool) // 所有接受的客户消息
	for {
		select {
		case msg := <-messages:
			// 所有接受的客户消息广播给所有的客户
			// 发送消息通道
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			// 监听客户的到来
			clients[cli] = true

		case cli := <-leaving:
			// 监听客户的离开
			delete(clients, cli)
			close(cli)
		}
	}
}

// handleConn 处理客户端每一个连接
func handleConn(conn net.Conn) {
	ch := make(chan string) // 对外发送消息通道
	go clientWriter(conn, ch) // 为每个客户创建一个消息写入的goroutine

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

// clientWriter 客户端消息写入
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		_, _ = fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func main() {
	// 主goroutine监听端口，接受客户端的网络连接
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// 对每一个网络连接创建一个新的goroutine
		go handleConn(conn)
	}
}
