//使用容量為1的channel確保最多一個goroutine存取共用變數
package main

var (
	sema    = make(chan struct{}, 1) //以二元號誌保護餘額
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} //取得token
	balance = balance + amount
	<-sema //釋放token
}

func Balance() int {
	sema <- struct{}{} //取得token
	b := balance
	<-sema //釋放token
	return b
}
