package reflect

import (
	"fmt"
	"reflect"
)

func ElemEntry() {

	//a := reflect.ValueOf(x)
	//b := a.Elem()
	//fmt.Println(b.CanSet())
	//fmt.Println(b.CanAddr())

	//byPoint()
	usingElem()
}

func usingElem() {
	x := 2
	fmt.Printf("x = %d\n", x)
	a := reflect.ValueOf(&x)
	b := a.Elem()
	b.SetInt(10)
	//p := b.Addr().Interface().(*int)
	//*p = 10
	fmt.Printf("x = %d\n", x)
}

func byPoint() {
	x := 2
	fmt.Printf("x = %d\n", x)
	a := reflect.ValueOf(&x)
	p := a.Interface().(*int)
	*p = 10
	fmt.Printf("x = %d\n", x)
}
