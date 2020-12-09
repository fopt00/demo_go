package goroutine

import "fmt"

func PipelineEntry() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 0; i < 101; i++ {
			naturals <- i
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}
}
