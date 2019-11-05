package main

import (
	"memotest"
	"sync"
)

//Func是要記憶的函式型別
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} //res就緒時關閉
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

type Memo struct {
	f     Func
	mu    sync.Mutex //保護快取
	cache map[string]*entry
}

//Get呼叫取得互斥鎖、查詢map中現有entry指標、未找到時分配並插入新的entry
func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		//此鍵的第一次請求
		//此goroutine負責計算值
		//並廣播就緒條件
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)

		close(e.ready) //散播就緒條件
	} else {
		//此鍵的重複請求
		memo.mu.Unlock()

		<-e.ready
	}
	return e.res.value, e.res.err
}

func main() {
	m := New(memotest.HTTPGetBody)

	//memotest.Concurrent(m) //平行
	memotest.Sequential(m)
}
