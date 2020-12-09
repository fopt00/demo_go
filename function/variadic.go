package function

import "fmt"

func VariadicEntry() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sum(a...))
}

func sum(values ...int) (sum int) {
	for _, val := range values {
		sum += val
	}
	return
}
