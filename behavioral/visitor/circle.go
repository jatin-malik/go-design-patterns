package visitor

type Circle struct {
	radius int
}

func (s *Circle) getType() string {
	return "Circle"
}

func (s *Circle) accept(v Visitor) {
	v.doForCircle(s)
}
