package structural

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
