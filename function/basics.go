package function

import "fmt"

func FunctionBasicEntry() {
	//sum:= add(1, 2)
	//fmt.Println(sum)

	//x, y := 1, 2
	//x, y = swap(x, y)
	//fmt.Println(x, y)

	//fmt.Println(namedAdd(3, 2))
	sumRecursiveEntry()
}

func sumRecursiveEntry() {
	fmt.Println(sumRecursive(100))
}

func sumRecursive(n int) int  {
	if n == 1 {
		return 1
	}
	return n + sumRecursive(n - 1)
}

func namedAdd(x, y int) (sum int) {
	sum = x + y
	return
}

func add(x, y int) int {
	return x + y
}

func swap(x, y int) (int, int) {
	return y, x
}
