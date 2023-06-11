package main

import "fmt"

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

func RunInterfacesDemo() {
	// car := Car{4}
	bike := Bike{2}

	if CheckCar(bike) {
		fmt.Println("Yes!")
	} else {
		fmt.Println("No!")
	}
}
