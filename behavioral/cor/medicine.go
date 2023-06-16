package cor

import "fmt"

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
