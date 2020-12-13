package reflect

import (
	"fmt"
	"reflect"
)

func IteratingStructEntry() {
	p := Ps{
		Name: "Allen",
		age:  19,
		Data: Data{
			19,
			20,
		},
	}

	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)
	fmt.Println("iterating fields")
	fmt.Println("-------------")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		value := v.Field(i)
		fmt.Printf("index:%d\nName:%s\nPkgPath:%s\nType:%v\nTag:%s\nOffset:%v\nIndex:%v\nAnonymous:%v\nValue:%v\n",
			i, f.Name, f.PkgPath, f.Type, f.Tag, f.Offset, f.Index, f.Anonymous, value)
		fmt.Println("-------------")
	}

	t = reflect.TypeOf(&p)

	fmt.Println("iterating methods")
	fmt.Println("-------------")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("index:%d\nName:%s\nPkgPath:%s\nType:%v\nFunc:%v\nIndex:%v\n", i, m.Name, m.PkgPath, m.Type, m.Func, m.Index)
		fmt.Println("-------------")
	}
}

type Data struct {
	height uint32
	width  uint32
}

type Ps struct {
	Name string `tips: this is name`
	age  int32 `how old are you?`
	Data
}

func (p *Ps) GetName() string {
	return p.Name
}

func (p *Ps) GetAge() int32 {
	return p.age
}
