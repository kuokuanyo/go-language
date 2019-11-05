//使快取具有並行安全，使用監視器同步化
//利用互斥鎖
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

//具有並行安全
func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	memo.mu.Unlock()
	return res.value, res.err
}

func main() {
	m := New(memotest.HTTPGetBody)

	memotest.Concurrent(m) //平行
	memotest.Sequential(m)
}
