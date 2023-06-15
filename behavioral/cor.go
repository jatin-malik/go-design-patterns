package behavioral

import "fmt"

// Chain of Responsibility is behavioral design pattern that allows passing request along the chain of potential
// handlers until one of them handles request.

// Hospital department are chain of handlers handling a patient (request)

type Patient struct {
	Name               string
	isRegistrationDone bool
	isCheckupDone      bool
	isMedicineGiven    bool
	isPaymentDone      bool
}

type Department interface {
	treat(*Patient)
	setNext(Department)
}

type Reception struct {
	next Department
}

func (r *Reception) treat(p *Patient) {
	if p.isRegistrationDone {
		fmt.Println("Already registered")
	} else {
		fmt.Println("Registering patient", p.Name)
		p.isRegistrationDone = true
	}

	if r.next != nil {
		r.next.treat(p)
	}

}

func (r *Reception) setNext(d Department) {
	r.next = d
}

type Doctor struct {
	next Department
}

func (d *Doctor) treat(p *Patient) {
	if p.isCheckupDone {
		fmt.Println("Already checkup done")
	} else {
		fmt.Println("Doing checkup of", p.Name)
		p.isCheckupDone = true
	}

	if d.next != nil {
		d.next.treat(p)
	}
}

func (d *Doctor) setNext(de Department) {
	d.next = de
}

type Medicine struct {
	next Department
}

func (m *Medicine) treat(p *Patient) {
	if p.isMedicineGiven {
		fmt.Println("Already given medicine")
	} else {
		fmt.Println("Giving medicine to patient", p.Name)
		p.isMedicineGiven = true
	}

	if m.next != nil {
		m.next.treat(p)
	}
}

func (m *Medicine) setNext(d Department) {
	m.next = d
}

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

func RunCORDemo() {
	mukesh := &Patient{
		Name:               "Mukesh",
		isRegistrationDone: true,
		isCheckupDone:      true,
	}
	reception := Reception{}
	doctor := Doctor{}
	reception.setNext(&doctor)

	medicine := Medicine{}
	doctor.setNext(&medicine)
	cashier := Cashier{}
	medicine.setNext(&cashier)

	// Mukesh goes to reception
	reception.treat(mukesh)

}
