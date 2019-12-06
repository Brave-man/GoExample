package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"goCode/review/channel/thumbnail"
)

// 同步执行
func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

// 异步执行 但是主goroutine不知道是否进行完异步处理，就结束了主goroutine, 导致文件没处理完
func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go thumbnail.ImageFile(f)
	}
}

func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		// 通过添加显示参数，可以确保当前go语句执行的是当前循环的f
		go func(f string) {
			thumbnail.ImageFile(f)
			ch <- struct{}{}
		}(f)
	}

	// 等待goroutine完成
	for range filenames {
		<-ch
	}
}

// 生成缩略图，如果任何步骤出错，它返回一个错误
func makeThumbnails4(filenames []string) error {
	errors := make(chan error)
	for _, f := range filenames {
		// 通过添加显示参数，可以确保当前go语句执行的是当前循环的f
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err
		}(f)
	}

	// 等待goroutine完成
	for range filenames {
		if err := <-errors; err != nil {
			return err // 错误做法，会导致goroutine泄漏
			// 当遇到第一个非nil的错误时，它将错误返回给调用者，调用者做了退出操作
			// 每一个现存的goroutine无法再向此通道发送消息，导致永久阻塞，永不终止
			// 这种情况会导致程序卡住或者系统内存耗尽
		}
	}
	return nil
}

func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}
	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}

// 生成缩略图并返回其文件占用的字节数
func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // 工作goroutine的个数
	for f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // 可以忽略错误
			sizes <- info.Size()
		}(f)
	}

	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}

func main() {
	photos := []string{
		"/Users/zhangzhiwei/go/src/goCode/review/channel/thumbnailUse/static/photo.jpeg",
		"/Users/zhangzhiwei/go/src/goCode/review/channel/thumbnailUse/static/photo1.jpg",
		"/Users/zhangzhiwei/go/src/goCode/review/channel/thumbnailUse/static/photo2.jpg",
		"/Users/zhangzhiwei/go/src/goCode/review/channel/thumbnailUse/static/photo3.jpg",
	}
	//now := time.Now()
	//_, _ = makeThumbnails5(photos)
	//fmt.Printf("%s\n", time.Since(now))

	now2 := time.Now()
	filenames := make(chan string, len(photos))
	for _, photo := range photos {
		filenames <- photo
	}
	close(filenames)
	fmt.Println(makeThumbnails6(filenames))
	fmt.Printf("%s\n", time.Since(now2))
}
