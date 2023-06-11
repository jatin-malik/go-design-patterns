package adapter

func RunDemo() {
	// We have a client who wants to connect his cable to the mac and windows computer
	client := Client{Name: "Elliot"}
	mac := &Mac{}
	windows := &Windows{}

	client.InsertIntoPort(mac)
	// But the windows computer does not have the same interface , it cannot accept cable connections
	// So we make a adapter around it , that can take cable connection , convert to serial and then connect to windows.
	windowsAdapter := &WindowsAdapter{Windows: windows}
	client.InsertIntoPort(windowsAdapter)
}
