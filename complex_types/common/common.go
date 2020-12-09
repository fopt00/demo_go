package common

type Person struct {
	Name string
	Age int
	hobby string
}

func CreatePerson(name string, age int, hobby string) Person {
	return Person{name, age, hobby}
}
