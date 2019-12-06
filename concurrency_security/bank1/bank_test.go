package bank1_test

import (
	"fmt"
	"testing"

	"goCode/review/concurrency_security/bank1"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	go func() {
		bank1.Deposit(200)
		fmt.Println("balance=", bank1.Balance())
		done <- struct{}{}
	}()

	go func() {
		bank1.Deposit(100)
		done <- struct{}{}
	}()

	<-done
	<-done

	if got, want := bank1.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
