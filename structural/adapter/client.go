package adapter

import "fmt"

type Client struct {
	Name string
}

// The client just wants to insert his cable into the USB-C port
func (c *Client) InsertIntoPort(com Computer) {
	fmt.Printf("%s has inserted his cable in the computer\n", c.Name)
	com.AcceptCableConnection()
}
