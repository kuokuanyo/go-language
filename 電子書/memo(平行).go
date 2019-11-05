package main

import (
	"memo"
	"memotest"
)

func main() {
	m := memo.New(memotest.HTTPGetBody)
	memotest.Concurrent(m) //平行

	memotest.Sequential(m)
}
