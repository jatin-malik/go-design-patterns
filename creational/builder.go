package creational

type VehicleProduct struct {
	Structure string
	Wheels    int
	Seats     int
}

type BuildProcess interface {
	// Every vehicle product builder must comply with these steps
	SetStructure() BuildProcess
	SetWheels() BuildProcess
	SetSeats() BuildProcess

	GetVehicle() VehicleProduct
}

// Can use singleton pattern for this director
type ManufacturingDirector struct {
	builder BuildProcess
}

func (m *ManufacturingDirector) SetBuilder(b BuildProcess) {
	m.builder = b
}

func (m *ManufacturingDirector) Build() {
	m.builder.SetStructure().SetSeats().SetWheels()
}

type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bike"
	return b
}

func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}

type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}
