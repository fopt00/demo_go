package complex_types

import "fmt"

func MapEntry() {
	//makeMap()
	//visitMap()
	deleteKeyFromMap()
}

func deleteKeyFromMap() {
	ages := map[string]int{
		"allen": 20,
		"luffy": 19,
		"zoro":  20,
		"sanji": 20,
	}
	delete(ages, "allen")
	for key, value := range ages {
		fmt.Println(key, value)
	}
}

func visitMap() {
	ages := map[string]int{
		"allen": 20,
		"luffy": 19,
		"zoro":  20,
		"sanji": 20,
	}
	value, exists := ages["allen1"]
	if exists {
		fmt.Println(value)
	}

	for k, v := range ages {
		fmt.Printf("%v : %d\n", k, v)
	}
}

func makeMap() {
	m1 := make(map[string]int)
	m2 := map[string]int{}
	m3 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	var m4 map[string]int
	fmt.Println(m1, m2, m3, m4)
}
