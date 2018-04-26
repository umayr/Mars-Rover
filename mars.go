package mars

import "fmt"

var (
	// Seq is a fixed array of possible directions in predefined sequence
	Seq = [4]string{"N", "E", "S", "W"}

	// ErrOutOfBound is an error returned if rover is moved out of the defined plateau
	ErrOutOfBound = fmt.Errorf("rover is out of bound")
	// ErrBadCommand is an error returned if the instruction set provided is not in the desired format
	ErrBadCommand = fmt.Errorf("command is invalid, only L, R, and M are valid ones")
)

const (
	// North represents north heading with respect to Seq array
	North = iota
	// East represents east heading with respect to Seq array
	East
	// South represents south heading with respect to Seq array
	South
	// West represents west heading with respect to Seq array
	West
)

const (
	// Left represents `L` command
	Left = "L"
	// Right represents `R` command
	Right = "R"
	// Move represents `M` command
	Move = "M"
)
