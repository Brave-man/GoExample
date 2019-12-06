package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func directs(dir string) []os.FileInfo {
	select {
	case tokens <- struct{}{}:
	case <-done:
		return nil
	}
	defer func() { <-tokens }() // 退出时 释放令牌
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	// walkDir在开始前轮训取消状态，如果已设置状态，什么都不做立即返回
	if cancelled() {
		return
	}

	for _, entry := range directs(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func printDiskUsage(nFiles, nBytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nFiles, float64(nBytes)/1e9)
}

func cancelled() bool {
	select {
	// 检测到输入后，关闭done通道，此处会从通道获取到零值struct{}{}
	case <-done:
		return true
	default:
		return false
	}
}

var verbose = flag.Bool("v", false, "show verbose progress messages")
var tokens = make(chan struct{}, 20)
var done = make(chan struct{}) // 取消通道

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// 当检测到输入时，取消遍历
	go func() {
		_, _ = os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	fileSizes := make(chan int64)
	var n sync.WaitGroup // 使用等待组
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes) // 为每一个目录创建一个goroutine
	}

	go func() {
		n.Wait() // 等待所有goroutine结束
		close(fileSizes) // 关闭文件通道
	}()

	// 定制输出结果
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nFiles, nBytes int64
loop:
	for {
		select {
		case <-done:
			// 耗尽fileSizes以允许已有的goroutine结束
			for range fileSizes {
				// 不执行任何操作
			}
			return
		case size, ok := <- fileSizes:
			if !ok {
				break loop
			}
			nFiles++
			nBytes += size
		case <-tick:
			printDiskUsage(nFiles, nBytes)
		}
	}
	printDiskUsage(nFiles, nBytes)
}
