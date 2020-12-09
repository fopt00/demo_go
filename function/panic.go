package function

import "fmt"

func PanicEntry() {
	//makeAPanic()
	testRecover()
}

func testRecover() {
	f2(3)
	fmt.Println("end")
}

func f2(x int) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("%v\n", p)
		}
	}()

	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("%d\n", x)
	f(x - 1)
}

func makeAPanic() {
	f(3)
	fmt.Println("end")
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("%d\n", x)
	f(x - 1)
}
