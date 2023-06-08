package creational

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
