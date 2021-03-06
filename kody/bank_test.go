package main

import (
	"sync"
	"testing"

	"bank"
)

func TestBank(t *testing.T) {
	// Wpłata [1..1000] równolegle.
	var n sync.WaitGroup
	for i := 1; i <= 1000; i++ {
		n.Add(1)
		go func(amount int) {
			bank.Deposit(amount)
			n.Done()
		}(i)
	}
	n.Wait()

	if got, want := bank.Balance(), (1000+1)*1000/2; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
