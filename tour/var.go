package tour

import "fmt"

var c, python, java bool
func Tester()  {
	var i int
	fmt.Println(i, c, python, java)
}

func Tester2()  {
	c, python, java = true, false, false
	fmt.Println(c, python, java)
}