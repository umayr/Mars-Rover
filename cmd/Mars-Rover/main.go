package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var maxRovers = 2

var (
	errInvalidEnv = fmt.Errorf("invalid value of environment variable MAX_ROVERS")
	errBadPlateau = fmt.Errorf("invalid input for plateau, it must be two coordinates separated with space")
	errBadLanding = fmt.Errorf("invalid input for rover landing, it must be two coordinates and heading (N, E, W, S) separated with spaces")
	errBadCommand = fmt.Errorf("invalid input for rover instructions, it must be a combination of L, R, and M")
)

func init() {
	m := os.Getenv("MAX_ROVERS")
	if m == "" {
		return
	}

	i, err := strconv.Atoi(m)
	if err != nil {
		fmt.Println(errInvalidEnv)
		os.Exit(1)
	}

	maxRovers = i
}

func input(msg string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	txt, err := reader.ReadString('\n')
	return strings.TrimSpace(txt), err
}

func main() {
Plateau:
	p, err := input("Plateau: ")
	if err != nil {
		fmt.Println(errBadPlateau)
		goto Plateau
	}

	b, err := parsePlateau(p)
	if err != nil {
		fmt.Println(err)
		goto Plateau
	}

	out := make([]string, maxRovers)

	for i := 0; i < maxRovers; i++ {
	Landing:
		p, err := input(fmt.Sprintf("Rover#%d Landing: ", i+1))
		if err != nil {
			fmt.Println(errBadLanding)
			goto Landing
		}

		r, err := parseRover(p)
		if err != nil {
			fmt.Println(err)
			goto Landing
		}

		r.Bound = *b

	Instruction:
		c, err := input(fmt.Sprintf("Rover#%d Instructions: ", i+1))
		if err != nil {
			fmt.Println(errBadCommand)
			goto Instruction
		}

		if err := r.Exec(c); err != nil {
			fmt.Println(errBadCommand)
			goto Instruction
		}

		out[i] = fmt.Sprint(r)
	}

	for i := 0; i < maxRovers; i++ {
		fmt.Printf("Rover#%d: %s\n", i+1, out[i])
	}
}
