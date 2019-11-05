//提供無並行安全的Func型別函式記憶化
package main

//快取呼叫Func的結果
type Memo struct { //大寫匯出
	f     Func
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

//method
//沒有並行安全
func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key] //res: result
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}

