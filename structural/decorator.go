package structural

import "fmt"

// Decorator is a structural pattern that allows adding new behaviors to objects dynamically
// by placing them inside special wrapper objects, called decorators.

// Let's take pizza as base product and toppings are layer above it

type Pizza interface {
	getPrice() int
	getName() string
}

// A concrete implementation of pizza
type FreshVeggie struct {
	Name string
}

func (p *FreshVeggie) getPrice() int {
	return 15
}

func (p *FreshVeggie) getName() string {
	return "FreshVeggie"
}

// Decorartors ( toppings ) wrapping this pizza

type CheeseTopping struct {
	p Pizza // Wraps the Pizza
}

func (c *CheeseTopping) getPrice() int {
	return 10 + c.p.getPrice()
}

func (c *CheeseTopping) getName() string {
	return "Cheese" + c.p.getName()
}

type PaneerTopping struct {
	p Pizza // Wraps the Pizza
}

func (t *PaneerTopping) getPrice() int {
	return 15 + t.p.getPrice()
}

func (c *PaneerTopping) getName() string {
	return "Paneer" + c.p.getName()
}

func RunDecoratorDemo() {
	// I want a pizza with cheese topping
	pizza := &FreshVeggie{}
	p := &CheeseTopping{pizza}
	fmt.Println("Got the pizza ", p.getName(), "with price", p.getPrice())
}
