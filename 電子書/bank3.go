//互斥鎖:sync.Mutex
//互斥鎖保護共用變數
//Lock可取得token
package main

import "sync"

var (
	mu      sync.Mutex //保護餘額
	balance int
)

func Deposit(amount int) {
	mu.Lock()                  //取得token
	balance = balance + amount //關鍵區(可自由讀取與修改共變數區域)
	mu.Unlock()                //釋放token
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}
