package main

import (
	"fmt"
	"log"
	"os"

	"goCode/review/channel/links"
)

// 令牌桶 确保并发请求限制在20个以内
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // 获取令牌
	list, err := links.Extract(url)
	<-tokens // 释放令牌
	if err != nil {
		log.Println(err)
	}
	return list
}

func main() {
	workList := make(chan []string)
	var n int // 等待发送到任务列表的数量
	// 从命令行参数开始
	n++
	go func() {
		workList <- os.Args[1:]
	}()

	// 并发爬取web
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-workList
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					workList <- crawl(link)
				}(link)
			}
		}
	}
}
