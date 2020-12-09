package function

import (
	"fmt"
	"time"
)

func DeferredEntry() {
	//simpleUse()
	//recordProcessTime()
	//recordParamsAndReturnValues()
	testDeferOrder()
}

func testDeferOrder() {
	defer func() {
		fmt.Println("one")
	}()

	defer func() {
		fmt.Println("two")
	}()

	defer func() {
		fmt.Println("three")
	}()
	fmt.Println("done")
}

func recordParamsAndReturnValues() {
	add2(1, 9)
}

func add2(x, y int) (z int) {
	defer func() {
		fmt.Printf("input: %d,%d output:%d\n", x, y, z)
	}()
	return x + y
}

func recordProcessTime() {
	defer trace("process")()
	time.Sleep(2 * time.Second)
}

func trace(s string) func() {
	start := time.Now()
	fmt.Printf("%s start\n", s)
	return func() {
		fmt.Printf("%s end. elsaps:%s\n", s, time.Since(start))
	}
}

func simpleUse() {
	fmt.Println("allocate some resource")
	defer free()
	fmt.Println("test alloc and free")
}

func free() {
	fmt.Println("free the resource")
}
