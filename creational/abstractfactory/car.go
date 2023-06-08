package abstractfactory

type Car interface {
	NumDoors() int
}

type LuxuryCar struct{}

func (c *LuxuryCar) NumWheels() int {
	return 4
}

func (c *LuxuryCar) NumSeats() int {
	return 5
}

func (c *LuxuryCar) NumDoors() int {
	return 4
}

type FamilyCar struct{}

func (c *FamilyCar) NumWheels() int {
	return 4
}

func (c *FamilyCar) NumSeats() int {
	return 5
}

func (c *FamilyCar) NumDoors() int {
	return 5
}
