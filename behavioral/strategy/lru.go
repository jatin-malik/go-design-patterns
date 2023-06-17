package strategy

import "fmt"

type LRUEviction struct {
}

func (e *LRUEviction) evict(c *Cache) {
	fmt.Println("Evicting the least recently used item")
}
