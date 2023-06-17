package strategy

/*
Strategy is a behavioral design pattern that turns a set of behaviors into objects and
makes them interchangeable inside original context object.

The original object, called context, holds a reference to a strategy object. The context delegates executing the
behavior to the linked strategy object. In order to change the way the context performs its work, other objects may
replace the currently linked strategy object with another one.

*/

func RunDemo() {

	lru := LRUEviction{}

	cache := InitCache(&lru, 3)

	cache.Add("name", "elliot")
	cache.Add("profession", "hacker")
	cache.Add("age", "25")

	cache.Add("enemies", "whiterose,dark army")

	fifo := FIFO{}
	cache.SetAlgo(&fifo)

	cache.Add("projects", "AllSafe")
}
