package goroutine

import (
	"fmt"
	"os"
	"time"
)

func SelectEntry() {
	tick := time.Tick(1 * time.Second)
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	for n := 10; n > 0; n-- {
		select {
		case <-tick:
			fmt.Println(n)
		case <-abort:
			fmt.Println("abort")
			return
		}
	}

	fmt.Println("launch")
}
