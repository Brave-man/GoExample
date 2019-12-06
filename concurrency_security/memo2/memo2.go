package memo2

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
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	memo.mu.Unlock()
	return res.value, res.err
}
