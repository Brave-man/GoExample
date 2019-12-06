package main

import (
	"fmt"
	"log"
	"os"

	"goCode/review/channel/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Println(err)
	}
	return list
}

func main() {
	workList := make(chan []string)  // 可能有重复url的列表
	unseenLinks := make(chan string) // 去重后的URL列表

	// 向任务列表中添加url参数
	go func() { workList <- os.Args[1:] }()

	// 创建20个goroutine
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { workList <- foundLinks }()
			}
		}()
	}

	// 主goroutine对url进行去重，并把没有爬去过的url发送给爬虫程序
	seen := make(map[string]bool)
	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
