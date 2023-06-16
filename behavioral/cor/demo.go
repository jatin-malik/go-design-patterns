package cor

// Chain of Responsibility is behavioral design pattern that allows passing request along the chain of potential
// handlers until one of them handles request.

// Hospital department are chain of handlers handling a patient (request)

func RunDemo() {
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
