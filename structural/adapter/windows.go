package adapter

import "fmt"

type Windows struct {
}

func (m *Windows) AcceptSerialConnection() {
	fmt.Println("Cable connected to the windows machine.")
}
