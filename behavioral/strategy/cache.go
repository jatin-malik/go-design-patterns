package strategy

// Let's have in-memory cache
type Cache struct {
	storage      map[string]string
	cap          int
	maxCap       int
	evictionAlgo EvictionAlgo
}

// You can add
func (c *Cache) Add(k, v string) {
	if c.cap == c.maxCap {
		c.Evict()
	}
	c.cap++
	c.storage[k] = v
}

func (c *Cache) Evict() {
	// Use the algo
	c.evictionAlgo.evict(c)
	c.cap--
}

func (c *Cache) SetAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func InitCache(e EvictionAlgo, cap int) *Cache {
	return &Cache{
		storage:      make(map[string]string),
		maxCap:       cap,
		evictionAlgo: e,
	}
}
