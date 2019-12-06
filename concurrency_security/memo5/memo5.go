package memo5

// Func 用于记忆函数的类型
type Func func(key string) (interface{}, error)

// result 是调用Func后的返回结果
type result struct {
	value interface{}
	err   error
}

// request 一个请求
type request struct {
	key      string
	response chan<- result
}

// Memo 缓存
type Memo struct {
	requests chan request
}

// New 构造一个缓存对象
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key: key, response: response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) // 调用f(key)
		}
		go e.deliver(req.response)
	}
}

type entry struct {
	res   result
	ready chan struct{} // res准备好后会被关闭
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	// 通知数据已经准备完毕
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// 等待数据准备完毕
	<-e.ready
	// 向客户端发送数据
	response <- e.res
}
