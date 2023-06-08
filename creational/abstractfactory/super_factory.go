package abstractfactory

import "errors"

const (
	CFactory = 1
	BFactory = 2
)

func GetFactory(f int) (VehicleFactory, error) {
	switch f {
	case CFactory:
		return &CarFactory{}, nil
	case BFactory:
		return &BikeFactory{}, nil
	default:
		return nil, errors.New("factory does not exist")
	}
}
