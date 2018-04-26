package main

import (
	"strconv"
	"strings"

	"github.com/umayr/Mars-Rover"
)

func parsePlateau(t string) (*mars.Point, error) {
	s := strings.Split(t, " ")
	if len(s) != 2 {
		return nil, errBadPlateau
	}

	sx, sy := s[0], s[1]

	x, err := strconv.Atoi(sx)
	if err != nil {
		return nil, errBadPlateau
	}

	y, err := strconv.Atoi(sy)
	if err != nil {
		return nil, errBadPlateau
	}

	return &mars.Point{X: uint(x), Y: uint(y)}, nil
}

func parseRover(t string) (*mars.Rover, error) {
	s := strings.Split(t, " ")
	if len(s) != 3 {
		return nil, errBadLanding
	}

	sx, sy, sh := s[0], s[1], s[2]

	x, err := strconv.Atoi(sx)
	if err != nil {
		return nil, errBadLanding
	}

	y, err := strconv.Atoi(sy)
	if err != nil {
		return nil, errBadLanding
	}

	var c int
	switch sh {
	case "N":
		c = mars.North
	case "E":
		c = mars.East
	case "S":
		c = mars.South
	case "W":
		c = mars.West
	default:
		return nil, errBadLanding
	}

	return &mars.Rover{
		Heading: mars.Heading{Cursor: c},
		Point:   mars.Point{X: uint(x), Y: uint(y)},
	}, nil
}
