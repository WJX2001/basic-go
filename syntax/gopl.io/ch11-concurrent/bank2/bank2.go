package bank2

var (
	sema    = make(chan struct{}, 1)
	balance int
)

func Deposit(amount int) {
	// 向sema里塞入数据
	sema <- struct{}{}
	balance = balance + amount
	<-sema // release token
}

func Balance() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}
