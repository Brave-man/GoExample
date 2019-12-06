package bank3_test

import (
	"sync"
	"testing"

	"goCode/review/concurrency_security/bank3"
)

func TestBank(t *testing.T) {
	balance := bank3.New()

	var n sync.WaitGroup
	for i := 1; i <= 1000; i++ {
		n.Add(1)
		go func(amount int) {
			balance.Deposit(amount)
			n.Done()
		}(i)
	}
	n.Wait()

	if got, want := balance.Balance(), (1+1000)*1000/2; got != want {
		t.Errorf("Balance=%d, want %d", got, want)
	}

	n.Add(2)
	go func() {
		balance.Withdraw(200)
		n.Done()
	}()

	go func() {
		balance.Withdraw(200)
		n.Done()
	}()
	n.Wait()

	if got, want := balance.Balance(), (1+1000)*1000/2-400; got != want {
		t.Errorf("Balance=%d, want %d", got, want)
	}
}
