package complex_types

import "fmt"

func ArrayEntry() {
	array := [...]string{0: "hello", 99: "world"}
	length := len(array)
	capacity := cap(array)
	fmt.Printf("%d, %d\n", length, capacity)

	// comparing
	arr1 := [...]int{1, 2, 3}
	arr2 := [...]int{1, 2, 3}
	fmt.Println("%T %T", arr1, arr2)
	fmt.Println(arr1 == arr2)

	// copy
	arr3 := [...]int{1, 2, 3}
	arr4 := arr3
	arr4[0] = 99
	fmt.Println(arr3, arr4)

	// call func using array as parameter
	testArray(arr3)
	fmt.Println("after call testArray", arr3)

	// pointer to an array
	arr5 := [...]int{1, 2, 3}
	p := &arr5
	firstElement := p[0]
	fmt.Println(firstElement)
	// modify element in an array by pointer
	p[0] = 0
	fmt.Println(arr5)

	// call func using pointer to an array
	arr6 := [...]int{1, 2, 3}
	testArrayUsingPointer(&arr6)
	fmt.Println("after call func using pointer to an array", arr6)
}

func testArrayUsingPointer(arr *[3]int) {
	// we can modify an array by using pointer
	arr[0] = 100
}

func testArray(arr [3]int) {
	fmt.Println("in the begin of testArray", arr)
	for i := range arr {
		arr[i] ++
	}
	fmt.Println("in the end of calling testArray", arr)
}
