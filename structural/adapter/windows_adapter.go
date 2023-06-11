package adapter

import "fmt"

type WindowsAdapter struct {
	Windows *Windows
}

func (w *WindowsAdapter) AcceptCableConnection() {
	fmt.Println("Converting cable connection into serial for windows machine")
	w.Windows.AcceptSerialConnection()
}
