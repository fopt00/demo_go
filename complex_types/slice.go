package complex_types

import (
	"fmt"
)

func SliceEntry() {
	//makeSlice()
	//testAppend()
	//makeSliceByArray()
	//emptySlice()
	insertionSortEntry()
}

func insertionSortEntry() {
	//s := []int{1, 2, 11, 3, 1}
	//var ret []int
	//ret = insertionSort(s)
	//fmt.Println(ret)

	a := [...]int{8, 6, 10, 7, 9, 3, 4, 5, 2, 1}

	fmt.Printf("a = %v\n", a)

	insertionSort(a[:5])
	insertionSort(a[5:])
	fmt.Println()

	fmt.Printf("a = %v\n", a)
}

func insertionSort(s []int) []int {
	for i := range s {
		e := s[i]
		j := i - 1
		for j >= 0 {
			if e < s[j] {
				s[j + 1] = s[j]
			} else {
				break
			}
			j --
		}
		s[j + 1] = e
	}
	return s
}

func emptySlice() {
	var s1 []int
	var s2 []int = nil
	s3 := []int(nil)
	s4 := []int{}
	fmt.Printf("%T:%[1]v, %t, len = %d, cap = %d\n", s1, s1 == nil, len(s1), cap(s1))  // []int:[], true, len = 0, cap = 0
	fmt.Printf("%T:%[1]v, %t, len = %d, cap = %d\n", s2, s2 == nil, len(s2), cap(s2))  // []int:[], true, len = 0, cap = 0
	fmt.Printf("%T:%[1]v, %t, len = %d, cap = %d\n", s3, s3 == nil, len(s3), cap(s3))  // []int:[], true, len = 0, cap = 0
	fmt.Printf("%T:%[1]v, %t, len = %d, cap = %d\n", s4, s4 == nil, len(s4), cap(s4))  // []int:[], false, len = 0, cap = 0

}

func makeSliceByArray() {
	arr := [...]int{1, 2, 3, 4, 5}
	slice1 := arr[2:3]
	slice2 := slice1[:3]
	fmt.Println(slice1, slice2)
}

func testAppend() {
	a := []int{1, 2, 3}
	_ = append(a, 4)
	fmt.Println(a)

	b := append(a, 4)
	fmt.Println(b)

	c := make([]int, 3, 4)
	c = append(c, 1)
	fmt.Println(c, cap(c), len(c))
}

func makeSlice() {
	slice1 := []int{1, 2, 3}
	fmt.Printf("%T, capacity = %d, length = %d\n", slice1, cap(slice1), len(slice1))

	// use make function to make a slice which has length of 10 and capacity of 10
	slice2 := make([]int, 10, 10)
	fmt.Printf("%T, capacity = %d, length = %d\n", slice2, cap(slice2), len(slice2))
}
