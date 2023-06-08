package abstractfactory

type Bike interface {
	Engine() string
}

type SportsBike struct{}

func (b *SportsBike) NumWheels() int {
	return 2
}

func (b *SportsBike) NumSeats() int {
	return 2
}

func (b *SportsBike) Engine() string {
	return "200CC"
}

type RegularBike struct{}

func (b *RegularBike) NumWheels() int {
	return 2
}

func (b *RegularBike) NumSeats() int {
	return 2
}

func (b *RegularBike) Engine() string {
	return "150CC"
}
