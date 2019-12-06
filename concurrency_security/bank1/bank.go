package bank1

var deposits = make(chan int) // 发送存款额
var balances = make(chan int) // 接受余额

func Deposit(amount int) { deposits <- amount }

func Balance() int { return <-balances }

func teller() {
	var balance int // balance 被限制在 teller goroutine 中
	for {
		select {
		case amount := <-deposits:
			// 只有在Deposit被调用后通道才会从阻塞状态打开，将值读出后，继续阻塞
			balance += amount
		case balances <- balance:
			// Balance在调用时，会阻塞，此时向balances通道发送值，Balance才会接受到值
		}
	}
}

func init() {
	go teller() // 启动监控goroutine
}
