package observer

import "fmt"

type Subject interface {
	subscribe(o Observer)
	unsubscribe(o Observer)
	notify()
}

// Concrete implementation

type Item struct {
	Name     string
	trackers map[int]Observer
	inStock  bool
}

func GetNewItem(name string) *Item {
	return &Item{Name: name, trackers: make(map[int]Observer)}
}

func (i *Item) updateStock() {
	fmt.Println("Item:", i.Name, "is back in stock.")
	i.inStock = true
	i.notify()
}

func (i *Item) subscribe(o Observer) {
	i.trackers[o.getId()] = o
}

func (i *Item) unsubscribe(o Observer) {
	delete(i.trackers, o.getId())
}

func (i *Item) notify() {
	for _, o := range i.trackers {
		o.notify(i.Name)
	}
}
