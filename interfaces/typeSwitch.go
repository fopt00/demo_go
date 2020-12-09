package interfaces

import (
	"fmt"
	"io"
	"os"
)

func TypeSwitchEntry() {
	show(5)                     // This is a number: 5
	show("hello world")         // This is a string: hello world
	show(true)                  // This is a boolean: true
	show(5.6)                   // This is a number: 5.6
	show(os.Stdout)             // This is a writer: *os.File
	show([]int{1, 2, 3, 4})   // I don't know what it is
}

func show(e interface{}) {
	switch v := e.(type) {
	case nil:
		fmt.Printf("Don't input nil\n")
	case int, uint, int32, uint32, float32, float64:
		fmt.Printf("This is a number: %v\n", v)
	case string:
		fmt.Printf("This is a string: %s\n", v)
	case bool:
		fmt.Printf("This is a boolean: %v\n", v)
	case io.Writer:
		fmt.Printf("This is a writer: %T\n", v)
	case *os.File:
		fmt.Printf("This is a *os.File\n")
	default:
		fmt.Printf("I don't know what it is\n")
	}
}
