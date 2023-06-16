package observer

// Observer is a behavioral design pattern that allows some objects to notify other objects about changes in their state.

// Let's say we have some customers who are interested in some items in our store and want to get notified
// whenever these items are in stock.

func RunDemo() {
	// Some customers
	cust1 := &Customer{1}
	cust2 := &Customer{2}
	cust3 := &Customer{3}

	// Some items
	xbox := GetNewItem("xbox")
	psp := GetNewItem("psp4")

	// Subscriptions
	xbox.subscribe(cust1)
	xbox.subscribe(cust2)

	psp.subscribe(cust2)
	psp.subscribe(cust3)

	// Stock inbound
	xbox.updateStock()
	psp.updateStock()

}
