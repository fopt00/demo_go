package concurrent

import (
	"fmt"
	"sync"
)

func RaceConditionEntry() {
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			addOne()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("the count is %d\n", count)
}

var count int
var m sync.Mutex

func addOne() {
	m.Lock()
	defer m.Unlock()
	count++
}
