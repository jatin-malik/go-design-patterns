package creational

import (
	"errors"
	"fmt"
)

// We have a business of a customized T-shirts
// We have some default tshirts ( prototypes ) which users can customize

type ShirtColor int

const (
	Blue  = 1
	Black = 2
	White = 3
)

type ShirtsCloner interface {
	GetClone(s int) (Shirt, error)
}

type Shirt struct {
	Color ShirtColor
	Price float32
	SKU   int
}

var WhiteShirtPrototype Shirt = Shirt{
	Color: White,
	Price: 15.00,
	SKU:   1,
}

var BlueShirtPrototype Shirt = Shirt{
	Color: Blue,
	Price: 16.00,
	SKU:   2,
}

var BlackShirtPrototype Shirt = Shirt{
	Color: Black,
	Price: 17.00,
	SKU:   3,
}

type ShirtsCache struct{}

func (s *ShirtsCache) GetClone(m int) (Shirt, error) {
	switch m {
	case Blue:
		// return cloned prototype of blue shirt
		return BlueShirtPrototype, nil
	case Black:
		return BlackShirtPrototype, nil
	case White:
		return WhiteShirtPrototype, nil
	default:
		return Shirt{}, errors.New("no such shirt prototype found")

	}
}

func GetShirtsCloner() (ShirtsCloner, error) {
	return &ShirtsCache{}, nil
}

func RunPrototypeDemo() {
	cloner, _ := GetShirtsCloner()
	proto1, _ := cloner.GetClone(Blue)
	proto1.SKU = 4
	proto2, _ := cloner.GetClone(Blue)
	fmt.Printf("%+v %p\n", proto1, &proto1)
	fmt.Printf("%+v %p", proto2, &proto2)
}
