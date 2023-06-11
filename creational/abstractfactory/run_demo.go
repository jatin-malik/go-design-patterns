package abstractfactory

import "fmt"

func RunDemo() {
	factory, err := GetFactory(CFactory)
	if err != nil {
		panic(err)
	}
	car, _ := factory.GetVehicle(LuxuryC)
	fmt.Println(car.(Car).NumDoors(), " ", car.NumSeats(), " ", car.NumWheels())
}
