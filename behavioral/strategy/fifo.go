package strategy

import "fmt"

type FIFO struct {
}

func (e *FIFO) evict(c *Cache) {
	fmt.Println("Evicting the first queued item")
}
