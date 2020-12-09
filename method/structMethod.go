package method

import (
	"fmt"
	"math"
)

func StructMethodEntry() {
	//p := Point{1, 2}
	//p.MoveTo(2, 3)
	//fmt.Println(p.Distance(Point{3, 5}))
	//p.Show()

	//c := Circle{Radius: 10.5}
	//c.Show()
	//c.MoveTo(1, 2)
	//c.Show()

	c := Circle{}
	d := Circle{}
	c.MoveTo(1, 2)

	dist := c.Distance(d.Point)
	fmt.Println(dist)
}

type Point struct {
	X, Y float64
}

type Circle struct {
	Radius float64
	Point
}

func (p *Point) MoveTo(x float64, y float64) {
	p.X, p.Y = x, y
}

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) Show() {
	fmt.Printf("Position(%v, %v)\n", p.X, p.Y)
}

func (c Circle) Show() {
	fmt.Printf("Position{%v, %v} radius:%v\n", c.X, c.Y, c.Radius)
}
