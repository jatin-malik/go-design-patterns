package behavioral

import "fmt"

// Memento is a behavioral design pattern that allows making snapshots of an object’s state and restoring it in future.

// The owner of the state creates the memento and they are stored in a seprate container class.

// lets say we have a text editor with a state.

type Editor struct {
	state string
}

type Memento struct {
	state string
}

func (m *Memento) getSavedState() string {
	return m.state
}

func (e *Editor) setState(state string) {
	e.state = state
}

func (e *Editor) getState() string {
	return e.state
}

func (e *Editor) createSnapshot() *Memento {
	return &Memento{
		state: e.state,
	}
}

func (e *Editor) restoreSnapshot(m *Memento) {
	e.state = m.getSavedState()
}

type History struct {
	mementos []*Memento
}

func GetNewHistory() *History {
	return &History{
		mementos: make([]*Memento, 0),
	}
}

func (h *History) addMemento(m *Memento) {
	h.mementos = append(h.mementos, m)
}

func (h *History) popMemento() *Memento {
	p := h.mementos[len(h.mementos)-1]
	h.mementos = h.mementos[:len(h.mementos)-1]
	return p
}

func (e *Editor) ShowEditorState() {
	fmt.Println("The editor says --> ", e.getState())
}

func RunMementoDemo() {
	// open a editor
	editor := &Editor{state: "Jot something down..."}
	editor.ShowEditorState()

	history := GetNewHistory()

	// user types something, we save the current state before updating.
	history.addMemento(editor.createSnapshot())
	editor.setState("This is my diary.")
	editor.ShowEditorState()

	history.addMemento(editor.createSnapshot())
	editor.setState("Today was good.")
	editor.ShowEditorState()

	history.addMemento(editor.createSnapshot())
	editor.setState("Beat seven kids, it was fun.")
	editor.ShowEditorState()

	// User hits undo
	editor.restoreSnapshot(history.popMemento())
	editor.ShowEditorState()

	// hits undo again
	editor.restoreSnapshot(history.popMemento())
	editor.ShowEditorState()

}
