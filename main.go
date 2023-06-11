package main

import "github.com/jatin-malik/go-design-patterns/structural/adapter"

func sum(a, b int) int {
	return a + b
}

func main() {
	// println("Hey guys!")
	// f := AddN(6)
	// fmt.Println(f(51))

	// // car := Car{4}
	// bike := Bike{2}

	// if CheckCar(bike) {
	// 	fmt.Println("Yes!")
	// } else {
	// 	fmt.Println("No!")
	// }

	// str := "what is this?"
	// fmt.Println(string(jsonify(str)))

	// obj := creational.GetInstance()
	// obj.AddOne()
	// obj.AddOne()
	// obj.AddOne()
	// obj1 := creational.GetInstance()
	// obj1.AddOne()
	// obj2 := creational.GetInstance()
	// x := obj2.AddOne()
	// fmt.Print(x)

	// b := &creational.CarBuilder{}
	// d := creational.ManufacturingDirector{}
	// d.SetBuilder(b)
	// d.Build()

	// v := b.GetVehicle()
	// v.Seats = 6

	// fmt.Printf("%+v", v)

	// pm, e := creational.GetPaymentMethod(2)
	// if e != nil {
	// 	panic(e)
	// }
	// fmt.Println(pm.Pay(25.12))
	// factory, err := abstractfactory.GetFactory(abstractfactory.CFactory)
	// if err != nil {
	// 	panic(err)
	// }
	// car, _ := factory.GetVehicle(abstractfactory.LuxuryC)
	// fmt.Println(car.(abstractfactory.Car).NumDoors(), " ", car.NumSeats(), " ", car.NumWheels())

	// cloner, _ := creational.GetShirtsCloner()
	// proto1, _ := cloner.GetClone(creational.Blue)
	// proto1.SKU = 4
	// proto2, _ := cloner.GetClone(creational.Blue)
	// fmt.Printf("%+v %p\n", proto1, &proto1)
	// fmt.Printf("%+v %p", proto2, &proto2)

	// son := structural.Son{}
	// son.P.Age = 32
	// son.P.Name = "Dave"
	// son.Nickname = "Charlie"

	// fmt.Println(son)

	// daughter := structural.Daughter{}
	// daughter.Name = "Dave"
	// daughter.Age = 32
	// daughter.Nickname = "Sweety"

	// fmt.Println(daughter)

	// We have a client who wants to connect his cable to the mac and windows computer
	client := adapter.Client{Name: "Elliot"}
	mac := &adapter.Mac{}
	windows := &adapter.Windows{}

	client.InsertIntoPort(mac)
	// But the windows computer does not have the same interface , it cannot accept cable connections
	// So we make a adapter around it , that can take cable connection , convert to serial and then connect to windows.
	windowsAdapter := &adapter.WindowsAdapter{Windows: windows}
	client.InsertIntoPort(windowsAdapter)

}
