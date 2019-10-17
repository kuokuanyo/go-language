//呼叫goroutine取得鎖(2次):一次用於查詢，一次用於查詢無結果時的更新
package main

import (
	"memotest"
	"sync"
)

type Memo struct {
	f     Func
	mu    sync.Mutex //保護快取
	cache map[string]result
}

//Func是要記憶的函式型別
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)

		//在兩個關鍵區段間，多個goroutine
		//可競爭計算f(key)並更新map
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
}

func main() {
	m := New(memotest.HTTPGetBody)

	memotest.Concurrent(m) //平行
	memotest.Sequential(m)
}
