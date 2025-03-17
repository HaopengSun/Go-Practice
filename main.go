package main

import (
	"fmt"
	"go-practice/math"
)

func main() {
	var s math.Shape

	s = math.Circle{Radius: 5}
	fmt.Println("Circle Area:", s.Area())
	fmt.Println("Circle Perimeter:", s.Perimeter())

	s = math.Rectangle{Width: 4, Height: 6}
	fmt.Println("Rectangle Area:", s.Area())
	fmt.Println("Rectangle Perimeter:", s.Perimeter())
}
