package cor

import "fmt"

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
