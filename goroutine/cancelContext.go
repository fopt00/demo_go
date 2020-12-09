package goroutine

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

func CancelContextEntry() {
	rand.Seed(time.Now().Unix())
	done := make(chan struct{})
	var wg sync.WaitGroup

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	wg.Add(3)

	go func() {
		fmt.Println(Get(done))
		wg.Done()
	}()

	go func() {
		fmt.Println(Get(done))
		wg.Done()
	}()

	go func() {
		fmt.Println(Get(done))
		wg.Done()
	}()

	wg.Wait()
}

func Get(done <-chan struct{}) string {
	duration := rand.Intn(5) + 2
	tick := time.After(time.Duration(duration) * time.Second)
	select {
	case <-tick:
		return fmt.Sprintf("get page %d", duration)
	case <-done:
		return fmt.Sprintf("cancel %d", duration)
	}
}
