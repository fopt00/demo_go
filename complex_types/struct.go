package complex_types

import (
	"demo_go/complex_types/common"
	"fmt"
)

func StructEntry() {
	//makeStruct()
	//makeStruct2()
	//visitStruct()
	//packageTest()
	nestedStruct()
}

func nestedStruct() {
	type Point struct {
		X, Y int
	}
	type Circle struct {
		// the center of circle
		p Point
		r float64
	}
	c := Circle{Point{1, 2}, 6.4}
	fmt.Println(c)
	c2 := Circle{p: Point{2, 3}, r:2.0}
	fmt.Println(c2)

	type CircleWithoutName struct {
		Point
		r float64
	}
	c3 := CircleWithoutName{Point{1, 2}, 3.1}
	fmt.Println(c3)
}

func packageTest() {
	allen := common.Person{
		Name: "Luffy",
		Age:  19,
	}
	fmt.Println(allen)

	luffy := common.CreatePerson("Luffy", 18, "art")
	fmt.Println(luffy)
}

type innerPerson struct {
	id     int
	name   string
	age    int
	height int
}

func visitStruct() {
	p := innerPerson{id: 1, name: "allen", age: 20, height: 180}
	fmt.Println(p.id)

	p2p := &p
	fmt.Println(p2p, *p2p)
	fmt.Println((*p2p).height)
	(*p2p).height = 190
	fmt.Println((*p2p).height)
}

func makeStruct2() {

	p := innerPerson{id: 1, name: "tom", age: 12}
	fmt.Println(p)
}

func makeStruct() {
	var p struct {
		id     int
		name   string
		age    int
		height int
	}

	p.name = "allen"
	p.age = 20
	p.height = 180
	fmt.Println(p)
}
