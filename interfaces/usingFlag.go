package interfaces

import (
	"flag"
	"fmt"
	"strings"
)

func UsingFlagEntry() {
	var sep = flag.String("s", " ", "separator")
	var n = flag.Bool("n", false, "omit trailing newline")
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}
}
