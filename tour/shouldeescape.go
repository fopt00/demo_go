package tour

import "fmt"

func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}
	return false
}

func TestLoop()  {
	Loop:
		for n := 0; n < 10; n ++ {
			switch n {
			case 1:
				doSth(n)
			case 2:
				break Loop
			default:
				fmt.Println(n)
			}
		}
		fmt.Println("after Loop")
}

func doSth(n int) {
	fmt.Println(n, "do sth.")
}