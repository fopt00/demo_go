package interfaces

import (
	"fmt"
	"syscall"
)

func ErrorEntry() {
	//err := errors.New("this is an error")
	//if err != nil {
	//	fmt.Println(err)
	//}

	_, err := syscall.Open("./test.txt", syscall.O_RDONLY, 0)

	if e, ok := err.(syscall.Errno); ok && e == 2 {
		fmt.Printf("Error: %v\n", e)
	}
}
