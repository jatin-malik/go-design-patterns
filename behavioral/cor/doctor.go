package cor

import "fmt"

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
