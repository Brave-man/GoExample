package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	_, _ = fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	_, _ = fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	_, _ = fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	for input.Scan() {
		// reverb1中直到上一次的回声枯竭在处理下一次，真实的回声会由三个独立的回声叠加组成，因此使用一个新的goroutine去异步处理
		// 当go语句执行的时候，计算echo函数对应的参数，因此input.Text()实在主goroutine中进行的
		go echo(c, input.Text(), 1*time.Second)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	log.Println("Listening on localhost:8000...")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println("Accept new conn...")
		// 完整处理一个请求，此处将请求放入一个新的goroutine中，做异步处理
		go handleConn(conn)
	}
}
