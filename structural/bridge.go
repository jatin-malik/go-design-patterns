package structural

import "fmt"

// Bridge is a structural design pattern that divides business logic or huge class
// into separate class hierarchies that can be developed independently.

// Computers and Printers

// Abstraction
type Computer interface {
	Print()
	SetPrinter(Printer)
}

// Implementation
type Printer interface {
	PrintFile()
}

type Mac struct {
	p Printer // Bridge between abstraction and implementation
}

func (m *Mac) Print() {
	fmt.Println("Print request coming from mac")
	m.p.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.p = p
}

type Windows struct {
	p Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request coming from windows")
	w.p.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.p = p
}

type HP struct {
}

func (h *HP) PrintFile() {
	fmt.Println("Printing file on HP printer")
}

type Epson struct {
}

func (e *Epson) PrintFile() {
	fmt.Println("Printing file on epson printer")
}

func RunBridgeDemo() {
	mac := &Mac{}
	windows := &Windows{}

	hp := &HP{}
	epson := &Epson{}

	mac.SetPrinter(hp)
	mac.Print()

	mac.SetPrinter(epson)
	mac.Print()

	windows.SetPrinter(hp)
	windows.Print()

	windows.SetPrinter(epson)
	windows.Print()
}
