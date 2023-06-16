package memento

func RunDemo() {
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
