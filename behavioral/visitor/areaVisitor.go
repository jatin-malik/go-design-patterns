package visitor

import (
	"fmt"
	"math"
)

type AreaVisitor struct {
	area float64
}

func (v *AreaVisitor) doForCircle(c *Circle) {
	fmt.Println("Calculating area for circle")
	pi := math.Pi
	t := c.radius * c.radius
	v.area = pi * float64(t)

}

func (v *AreaVisitor) doForSquare(s *Square) {
	fmt.Println("Calculating area for square")
	t := float64(s.side)
	v.area = t * t
}

func (v *AreaVisitor) doForRectangle(r *Rectangle) {
	fmt.Println("Calculating area for rectangle")
	v.area = float64(r.l) * float64(r.b)
}
