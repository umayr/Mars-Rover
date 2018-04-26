package mars

import (
	"fmt"
	"strings"
)

// Rover is a struct that represents a rover along with its current location, heading and boundaries
type Rover struct {
	Heading
	Point

	Bound Point
}

// String lets Rover struct implement Stringer interface to display its values in a string format
func (r Rover) String() string {
	return fmt.Sprintf("%d %d %s", r.X, r.Y, Seq[r.Cursor])
}

// Move moves a rover single cell ahead in current direction and returns an error if it gets out of bound
// In case of error, it reverts the value of X and Y for a rover
func (r *Rover) Move() error {
	x, y := r.X, r.Y
	switch r.Cursor {
	case North:
		r.Y++
	case East:
		r.X++
	case South:
		r.Y--
	case West:
		r.X--
	}

	if r.Out(&r.Bound) {
		r.X, r.Y = x, y
		return ErrOutOfBound
	}

	return nil
}

// Exec takes a set of instructions and executes them. It returns an error if either its a bad instruction
// set (any command other than L, R and M) or it is out of bound after a move command
func (r *Rover) Exec(q string) error {
	for _, v := range strings.Split(q, "") {
		switch v {
		case Left:
			r.Left()
		case Right:
			r.Right()
		case Move:
			if err := r.Move(); err != nil {
				return err
			}
		default:
			return ErrBadCommand
		}
	}
	return nil
}
