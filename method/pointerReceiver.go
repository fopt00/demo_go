package method

func PointerReceiverEntry() {
	allen := Person{"allen", 20}
	allen.sayHello()
	allen.grow()
	allen.sayHello()

	allenPtr := new(Person)
	allenPtr.grow()
	allenPtr.sayHello()
}

func (p *Person) grow() {
	p.Age++
}
