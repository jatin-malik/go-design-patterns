package abstractfactory

import "errors"

type VehicleFactory interface {
	GetVehicle(t int) (Vehicle, error)
}

const (
	LuxuryC  = 1
	FamilyC  = 2
	SportsB  = 3
	RegularB = 4
)

type CarFactory struct {
}

func (c *CarFactory) GetVehicle(t int) (Vehicle, error) {
	switch t {
	case LuxuryC:
		// Luxury car
		return &LuxuryCar{}, nil
	case FamilyC:
		// Family Car
		return &FamilyCar{}, nil
	default:
		return nil, errors.New("car type does not exist")
	}
}

type BikeFactory struct {
}

func (c *BikeFactory) GetVehicle(t int) (Vehicle, error) {
	switch t {
	case SportsB:
		// Sports bike
		return &SportsBike{}, nil
	case RegularB:
		// Regular Car
		return &RegularBike{}, nil
	default:
		return nil, errors.New("bike type does not exist")
	}
}
