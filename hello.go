package main

import (
	"fmt"
)

type (
	Shape interface {
		Area() float64
	}
	Rectangle struct {
		Base, Height float64
	}

	Square struct {
		Side float64
	}
)

func (r *Rectangle) Area() float64 {
	return r.Base * r.Height
}

func (s *Square) Area() float64 {
	return s.Side * s.Side
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle: {base: %v, height: %v}", r.Base, r.Height)
}

func main() {
	fmt.Println("Hello, World!")
	rectangle := Rectangle{4.2, 8.0}
	fmt.Println(rectangle)
	fmt.Println(rectangle.Area())

	square := Square{4.3}
	fmt.Println(square)
	fmt.Println(square.Area())

	myFirstMap := map[string]Rectangle{
		"r1": {3, 4},
		"r2": {4, 2},
	}

	fmt.Println(myFirstMap)

	var someShape Shape
	someShape = &rectangle
	fmt.Println(someShape)
}
