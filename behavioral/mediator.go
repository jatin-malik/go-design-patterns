package behavioral

import "fmt"

// Mediator is a behavioral design pattern that reduces coupling between components of a program
// by making them communicate indirectly, through a special mediator object

// In a railway station , station manager is the mediator , every train talks to him for permissions

type Train interface {
	arrive()
	depart()
}

type TrainMediator interface {
	CanArrive(Train) bool
	NotifyDeparture()
}

// concrete implementation of mediator

type StationManager struct {
	isPlatformFree bool
	queue          []Train
}

func GetNewStationManager() *StationManager {
	return &StationManager{
		isPlatformFree: true,
		queue:          make([]Train, 0),
	}
}

func (sm *StationManager) CanArrive(t Train) bool {
	if sm.isPlatformFree {
		sm.isPlatformFree = false
		return true
	}
	// add the train in the queue
	sm.queue = append(sm.queue, t)
	return false
}

func (sm *StationManager) NotifyDeparture() {
	// Since a traind departed, platform is free now
	if !sm.isPlatformFree {
		sm.isPlatformFree = true
	}
	// Tell the first train in queue to come to platform
	t := sm.queue[0]
	sm.queue = sm.queue[1:]
	fmt.Println("SM - Asking waiting train to come on platform")
	t.arrive()
}

// Concrete implementation of train

type HPExpress struct {
	SM TrainMediator
}

func (t *HPExpress) arrive() {
	// Check if it can arrive
	if t.SM.CanArrive(t) {
		fmt.Println("Himachal Express arrving on platform.")
	} else {
		fmt.Println("Himachal Express - waiting for the platform to clear")
	}
}

func (t *HPExpress) depart() {
	fmt.Println("Himachal Express departing from platform")
	t.SM.NotifyDeparture()
}

type PaschimExpress struct {
	SM TrainMediator // to talk with the station manager
}

func (t *PaschimExpress) arrive() {
	// Check if it can arrive
	if t.SM.CanArrive(t) {
		fmt.Println("Paschim Express arrving on platform.")
	} else {
		fmt.Println("Paschim Express - waiting for the platform to clear")
	}
}

func (t *PaschimExpress) depart() {
	fmt.Println("Paschim Express departing from platform")
	t.SM.NotifyDeparture()
}

func RunMediatorDemo() {

	sm := GetNewStationManager()

	// make two trains
	he := &HPExpress{SM: sm}
	pe := &PaschimExpress{SM: sm}

	he.arrive()
	pe.arrive()
	he.depart()
}
