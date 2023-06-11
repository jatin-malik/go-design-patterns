package adapter

import "fmt"

type Mac struct {
}

func (m *Mac) AcceptCableConnection() {
	fmt.Println("Cable connected to the mac machine.")
}
