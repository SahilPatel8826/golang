package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	breadth float64
	length  float64
}
type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.breadth * r.length
}
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	fmt.Println("Area of Rectangle:", Rectangle{breadth: 5, length: 10}.Area())
	fmt.Println("Area of Circle:", Circle{radius: 3}.Area())
}
