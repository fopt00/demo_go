package reflect

import (
	"fmt"
	"reflect"
)

func BasicEntry() {
	var i int32 = 3
	var f float64 = 3.12
	t1 := reflect.TypeOf(i)
	t2 := reflect.TypeOf(f)

	v1 := reflect.ValueOf(i)
	v2 := reflect.ValueOf(f)

	fmt.Printf("typeof i = %v\n", t1.String())
	fmt.Printf("typeof j = %v\n", t2.String())

	fmt.Printf("valueof i = %v\n", v1)
	fmt.Printf("valueof j = %v\n", v2)
}
