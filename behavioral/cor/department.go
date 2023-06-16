package cor

type Department interface {
	treat(*Patient)
	setNext(Department)
}
