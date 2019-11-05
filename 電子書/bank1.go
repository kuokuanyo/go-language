package bank

var deposits = make(chan int) //發送存款金額
var balances = make(chan int) //接收餘額

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
	var balance int //餘額受teller限制
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() //啟動監視器goroutine
}
