package main

import "fmt"

func main() {
	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle area: %f\n", calcArea(circle))
	fmt.Printf("Rectangle area: %f\n", calcArea(rectangle))
}

//go run Shape.go Interface.go
