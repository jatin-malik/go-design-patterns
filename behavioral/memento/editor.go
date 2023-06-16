package memento

import "fmt"

type Editor struct {
	state string
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

func (e *Editor) ShowEditorState() {
	fmt.Println("The editor says --> ", e.getState())
}
