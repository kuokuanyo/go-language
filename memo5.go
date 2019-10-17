package main

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

//Func套用鍵的請求訊息
type request struct {
	key      string
	response chan<- result //用戶端要求結果
}

type Memo struct{ requests chan request }

//回傳f，之後用戶端必須呼叫Close
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			//第一次請求此鍵
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) //呼叫f(key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	//評估函式
	e.res.value, e.res.err = f(key)
	//廣播就緒條件
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	//等待就緒條件
	<-e.ready
	//發送結果給用戶端
	response <- e.res
}
