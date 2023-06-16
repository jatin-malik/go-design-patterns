package observer

import "fmt"

type Observer interface {
	getId() int
	notify(string)
}

type Customer struct {
	Id int
}

func (c *Customer) getId() int {
	return c.Id
}

func (c *Customer) notify(item string) {
	fmt.Println("Notifying customer", c.Id, "about item", item)
}
