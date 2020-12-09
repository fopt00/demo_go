package interfaces

import "fmt"

func BasicEntry() {
	var d Dog
	var c Cat

	AnimalCall(&d)
	AnimalCall(c)
}

type Caller interface {
	Call() string
}

type Dog struct {

}

type Cat struct {

}

func (d *Dog) Call() string  {
	return "Wang wang..."
}

func (c Cat) Call() string  {
	return "Meow meow..."
}

func AnimalCall(caller Caller)  {
	fmt.Printf("This animal call: %s\n", caller.Call())
}