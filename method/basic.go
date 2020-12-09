package method

import "fmt"

type Names []string
type Second int

func (n Names) getLen() int {
	return len(n)
}

func (s Second) tell() {
	fmt.Printf("Now the second is %d", s)
}

func BasicEntry() {
	//allen := Person{Name: "allen", Age: 20}
	//allen.sayHello()
	songNames := Names{"Without you", "Friendly Pressure"}
	fmt.Println(songNames.getLen())

	s := Second(10)
	s.tell()
}

type Person struct {
	Name string
	Age  int
}

func (p Person) sayHello() {
	fmt.Printf("Hello, my name is %s and my age is %d\n", p.Name, p.Age)
}
