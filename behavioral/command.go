package behavioral

import "fmt"

// Command is behavioral design pattern that converts requests or simple operations into objects.
// The conversion allows deferred or remote execution of commands, storing command history, etc.

// We have a TV that can be turned on/off.

type Device interface {
	on()
	off()
}

type TV struct {
	isRunning bool
}

func (t *TV) on() {
	t.isRunning = true
	fmt.Println("Turning TV on")
}

func (t *TV) off() {
	t.isRunning = false
	fmt.Println("Turning TV off")
}

// Two ways of interacting with TV, TV buttons or remote buttons.

// Each button ( any button ) when pressed , executes a command. So lets embed a command in each button.

type Command interface {
	execute()
}

type OnCommand struct {
	d Device
}

func (c *OnCommand) execute() {
	// turn device on
	c.d.on()
}

type OffCommand struct {
	d Device
}

func (c *OffCommand) execute() {
	// turn device off
	c.d.off()
}

type Button interface {
	press()
}

type RemoteButton struct {
	c Command
}

func (b *RemoteButton) press() {
	b.c.execute()
}

func RunCommandDemo() {
	// Make a TV
	tv := &TV{}

	// Make buttons
	onButton := &RemoteButton{}
	offButton := &RemoteButton{}

	// embed commands
	onButton.c = &OnCommand{tv}
	offButton.c = &OffCommand{tv}

	// Interact
	onButton.press()

	fmt.Println("Bored now.")

	offButton.press()
}
