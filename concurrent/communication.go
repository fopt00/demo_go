package concurrent

import (
	"fmt"
	"sync"
	"time"
)

var deposits = make(chan int)
var balances = make(chan int)

func CommunicationEntry() {
	go monitor()
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for n := 0; n < 10; n++ {
			DoDeposit(100)
			time.Sleep(1 * time.Second)
		}
		wg.Done()
	}()

	go func() {
		for n := 0; n < 10; n++ {
			DoDeposit(100)
			time.Sleep(1 * time.Second)
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Printf("balance is %d\n", DoBalance())
}

func monitor() {
	var balance int

	for {
		select {
		case amount := <-deposits:
			balance += amount
			fmt.Printf("now the balance is %d\n", balance)
		case balances <- balance:
		}
	}
}

func DoDeposit(amount int) {
	deposits <- amount
}

func DoBalance() int {
	return <-balances
}