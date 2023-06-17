package visitor

import "fmt"

// Visitor is a behavioral design pattern that allows adding new behaviors to existing class hierarchy
// without altering any existing code.

func RunDemo() {
	circle := &Circle{radius: 3}
	square := &Square{side: 5}
	rect := &Rectangle{l: 4, b: 6}

	// Calculate area for each class
	visitor := &AreaVisitor{}

	circle.accept(visitor)
	fmt.Println("The area for cirlce is", visitor.area)
	square.accept(visitor)
	fmt.Println("The area for square is", visitor.area)
	rect.accept(visitor)
	fmt.Println("The area for rectangle is", visitor.area)
}
