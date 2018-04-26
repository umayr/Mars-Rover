package mars

import "fmt"

// Point is a struct that denotes a point with X and Y coordinate
type Point struct {
	X, Y uint
}

// String method lets Point struct implement Stringer interface to display Point values in string format
func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

// Out takes pointer to another Point and determines if the instance is out of bounds
func (p *Point) Out(q *Point) bool {
	return p.X > q.X || p.Y > q.Y
}
