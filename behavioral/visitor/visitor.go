package visitor

type Visitor interface {
	doForCircle(*Circle)
	doForSquare(*Square)
	doForRectangle(*Rectangle)
}
