package memo4

import "sync"

// Func 用于记忆函数的类型
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err error
}

type entry struct {
	res result
	ready chan struct{} // res准备好后会被关闭
}

type Memo struct {
	f Func
	mu sync.Mutex
	cache map[string]*entry
}

// New 构造一个缓存对象
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// 对key的第一次访问，这个goroutine负责计算数据和广播数据
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)

		close(e.ready) // 广播数据已准备完毕的消息
	} else {
		// 对这个key重复访问
		memo.mu.Unlock()

		<-e.ready // 等待数据准备完毕,这个操作一直会阻塞到通道关闭为止
	}
	return e.res.value, e.res.err
}
