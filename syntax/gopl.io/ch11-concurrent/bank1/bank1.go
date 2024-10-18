package main

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

type withdrawChan struct {
	ch     chan bool
	amount int
}

var withdraws = make(chan withdrawChan)

func Withdraw(amount int) bool {
	withDr := withdrawChan{
		amount: amount,
		ch:     make(chan bool),
	}

	withdraws <- withDr
	return <-withDr.ch
}

// 监听 goroutin
func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case withDr := <-withdraws:
			amount := withDr.amount
			if amount > balance {
				withDr.ch <- false
			} else {
				balance -= amount
				withDr.ch <- true
			}
		}
	}
}

func init() {
	go teller()
}
