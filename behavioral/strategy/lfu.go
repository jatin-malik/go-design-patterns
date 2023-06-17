package strategy

import "fmt"

type LFUEviction struct {
}

func (e *LFUEviction) evict(c *Cache) {
	fmt.Println("Evicting the least frequently used item")
}
