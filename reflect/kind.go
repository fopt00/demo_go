package reflect

import (
	"fmt"
	"reflect"
)

func KindOfEntry()  {
	var s Second = 100
	t := reflect.TypeOf(s)
	fmt.Printf("s.Type.String: %v\n", t.String())
	fmt.Printf("s.Type: %v\n", t)

	k := t.Kind()
	fmt.Printf("s.Type.kind: %d:%v\n", k, k)
	fmt.Println()

	v := reflect.ValueOf(s)
	fmt.Printf("s.Value.String: %v\n", v.String())
	fmt.Printf("s.Value: %v\n", v)
	fmt.Println()

	p := Person{Name: "jack", Age: 12}
	t = reflect.TypeOf(p)
	fmt.Printf("p.Type.String: %v\n", t.String())
	fmt.Printf("p.Type: %v\n", t)
	k = t.Kind()
	fmt.Printf("p.Type.kind: %d:%v\n", k, k)
	fmt.Println()
	v = reflect.ValueOf(p)
	fmt.Printf("p.Value.String: %v\n", v.String())
	fmt.Printf("p.Value: %v\n", v)
}

type Second int
type Person struct {
	Name string
	Age int
}