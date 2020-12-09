package concurrent

import (
	"fmt"
	"sync"
)

func SyncOnceEntry() {
	//tryOnce()
	//tryOnceWithGoRoutine()
	singleton()
}

type Configure struct {
	volume int32
}

var once sync.Once
var configure *Configure

func initConfigure() {
	configure = &Configure{10}
}

func GetInstance() *Configure {
	once.Do(initConfigure)
	return configure
}

func singleton() {
	var conf *Configure
	fmt.Printf("%p\n", conf)

	var wg sync.WaitGroup

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			conf = GetInstance()
			fmt.Printf("%p\n", conf)
			wg.Done()
		}()
	}

	wg.Wait()
}

func hello() {
	fmt.Println("hello")
}
func tryOnceWithGoRoutine() {
	var once sync.Once
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		go func() {
			once.Do(hello)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("first")
	fmt.Println("second")
}

func tryOnce() {
	var once sync.Once
	once.Do(hello)
	fmt.Println("first")
	once.Do(hello)
	fmt.Println("second")
}
