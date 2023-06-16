package memento

// Memento is a behavioral design pattern that allows making snapshots of an object’s state and restoring it in future.

// The owner of the state creates the memento and they are stored in a seprate container class.

// lets say we have a text editor with a state.

type Memento struct {
	state string
}

func (m *Memento) getSavedState() string {
	return m.state
}
