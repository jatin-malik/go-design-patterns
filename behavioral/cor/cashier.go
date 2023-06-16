package cor

import "fmt"

type Cashier struct {
	next Department
}

func (c *Cashier) treat(p *Patient) {
	if p.isPaymentDone {
		fmt.Println("Already paid")
	} else {
		fmt.Println("Payment being done by patient", p.Name)
		p.isPaymentDone = true
	}

	if c.next != nil {
		c.next.treat(p)
	}
}

func (c *Cashier) setNext(d Department) {
	c.next = d
}
