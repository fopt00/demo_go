package concurrent

import (
	"fmt"
	"sync"
)

var balance int
var mu sync.Mutex

func MutexEntry() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		alice(100)
		wg.Done()
	}()

	go func() {
		bob(200)
		wg.Done()
	}()

	wg.Wait()
	fmt.Printf("balance: %d\n", Balance())
}

func alice(amount int) {
	fmt.Printf("Alice deposit %d\n", amount)
	Deposit(amount)
}

func bob(amount int) {
	fmt.Printf("Bob deposit %d\n", amount)
	Deposit(amount)
}

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	balance += amount
}

func Balance() int {
	mu.Lock()
	mu.Unlock()
	return balance
}
