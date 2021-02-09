package tour

import (
	"math/rand"
	"time"
)

func Add(x, y int) int {
	return x + y
}

func Swap(x, y int) (int, int) {
	return y, x
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Split(sum int) (x, y int) {
	x = sum * rand.Intn(10) + 10
	y = sum - x
	return
}
