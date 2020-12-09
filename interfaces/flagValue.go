package interfaces

import (
	"flag"
	"fmt"
)

func FlagValueEntry() {
	p := Person{}
	flag.Var(&p, "p", "set the Person")
	flag.Parse()
	fmt.Printf("%v\n", &p)
}

type Person struct {
	Name string
	Age int
}

func (p *Person) Set(s string) error  {
	var name string
	var age int
	_, err := fmt.Sscanf(s, "%s %d", &name, &age)
	*p = Person{name, age}
	return err
}

func (p *Person) String() string  {
	return fmt.Sprintf("name:%s age:%d", p.Name, p.Age)
}
