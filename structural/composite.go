package structural

import "fmt"

type Parent struct {
	Name string
	Age  int
}

type Son struct {
	// I want to inherit from Parent but as Go has composition over Inheritance
	P        Parent // Direct composition
	Nickname string
}

type Daughter struct {
	// I want to inherit from Parent but as Go has composition over Inheritance
	Parent   // Embedded composition
	Nickname string
}

func RunCompositeDemo() {
	son := Son{}
	son.P.Age = 32
	son.P.Name = "Dave"
	son.Nickname = "Charlie"

	fmt.Println(son)

	daughter := Daughter{}
	daughter.Name = "Dave"
	daughter.Age = 32
	daughter.Nickname = "Sweety"

	fmt.Println(daughter)
}
