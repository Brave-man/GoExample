package memo3

import "sync"

// Func 用于记忆函数的类型
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err error
}

type Memo struct {
	f Func
	mu sync.Mutex
	cache map[string]result
}

// New 构造一个缓存对象
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)
		// 在两个临界区域之前，可能会有多个goroutine来计算f(key)，并且更新map
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
}
