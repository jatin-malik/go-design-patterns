package memento

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
