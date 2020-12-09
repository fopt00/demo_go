package function

import "fmt"

func AdvancedEntry() {
	//functionValue()
	//anonymousFunc()
	statefulEntry()
}

func statefulEntry() {
	addOne := step(0)
	fmt.Println(addOne())
	fmt.Println(addOne())
	fmt.Println(addOne())
	fmt.Println(addOne())
	fmt.Println(addOne())
	fmt.Println(addOne())
}

func step(x int) func() int {
	return func() int {
		x++
		return x
	}
}

func anonymousFunc() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(2, 3))
}

func functionValue() {
	val := testAdd
	var f func(int, int) int
	f = val
	fmt.Printf("%T\n", f)
	fmt.Println(f(1, 2))
	f = testSub
	fmt.Printf("%T\n", f)
	fmt.Println(f(1, 2))
}

func testAdd(x, y int) int {
	return x + y
}

func testSub(x, y int) int {
	return x - y
}
