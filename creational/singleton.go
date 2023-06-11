package creational

import "fmt"

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

func RunSingletonDemo() {
	obj := GetInstance()
	obj.AddOne()
	obj.AddOne()
	obj.AddOne()
	obj1 := GetInstance()
	obj1.AddOne()
	obj2 := GetInstance()
	x := obj2.AddOne()
	fmt.Print(x)
}
