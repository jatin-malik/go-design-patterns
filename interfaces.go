package main

type Vehicle interface {
	giveWheels() int
}

type Car struct {
	wheels int
}

type Bike struct {
	wheels int
}

func (c Car) giveWheels() int {
	return c.wheels
}

func (b Bike) giveWheels() int {
	return b.wheels
}

// car checker
func CheckCar(v Vehicle) bool {
	return v.giveWheels() == 4
}
