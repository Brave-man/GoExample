package bank3

import "sync"

type Balance struct {
	sync.Mutex
	balance int
	ws sync.Mutex // 提现专用互斥锁
}

// Deposit 存款操作
func (b *Balance) Deposit(amount int) {
	b.Lock()
	b.deposit(amount)
	b.Unlock()
}

// Withdraw 提现操作
func (b *Balance) Withdraw(amount int) bool {
	b.ws.Lock()
	defer b.ws.Unlock()
	b.Deposit(-amount)
	if b.Balance() < 0 {
		b.Deposit(amount)
		return false
	}
	return true
}

// 这个方法要求已经获取互斥锁
func (b *Balance) deposit(amount int) {
	b.balance += amount
}

// Balance 余额操作
func (b *Balance) Balance() int {
	b.Lock()
	defer b.Unlock()
	return b.balance
}

// New 获取一个新对象
func New() *Balance {
	return &Balance{}
}
