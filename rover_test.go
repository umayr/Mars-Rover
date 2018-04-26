package mars

import "testing"

func TestMove(t *testing.T) {
	r0 := &Rover{
		Heading{North},
		Point{0, 0},
		Point{2, 2},
	}

	r0.Move()
	if r0.Y != 1 {
		t.Errorf("rover should be at (0, 1) instead of %s", r0.Point)
	}

	r0.Move()
	if r0.Y != 2 {
		t.Errorf("rover should be at (0, 2) instead of %s", r0.Point)
	}

	err := r0.Move()
	if err == nil {
		t.Errorf("rover should be out of bound now instead of being at %s", r0.Point)
	}

	r1 := &Rover{
		Heading{East},
		Point{0, 2},
		Point{2, 2},
	}

	r1.Move()
	if r1.X != 1 {
		t.Errorf("rover should be at (1, 2) instead of %s", r1.Point)
	}

	r1.Move()
	if r1.X != 2 {
		t.Errorf("rover should be at (2, 2) instead of %s", r1.Point)
	}

	err = r1.Move()
	if err == nil {
		t.Errorf("rover should be out of bound now instead of being at %s", r1.Point)
	}

	r2 := &Rover{
		Heading{South},
		Point{2, 2},
		Point{2, 2},
	}

	r2.Move()
	if r2.Y != 1 {
		t.Errorf("rover should be at (2, 1) instead of %s", r2.Point)
	}

	r2.Move()
	if r2.Y != 0 {
		t.Errorf("rover should be at (2, 0) instead of %s", r2.Point)
	}

	err = r2.Move()
	if err == nil {
		t.Errorf("rover should be out of bound now instead of being at %s", r2.Point)
	}

	r3 := &Rover{
		Heading{West},
		Point{2, 2},
		Point{2, 2},
	}

	r3.Move()
	if r3.X != 1 {
		t.Errorf("rover should be at (1, 2) instead of %s", r3.Point)
	}

	r3.Move()
	if r3.X != 0 {
		t.Errorf("rover should be at (0, 2) instead of %s", r3.Point)
	}

	err = r3.Move()
	if err == nil {
		t.Errorf("rover should be out of bound now instead of being at %s", r3.Point)
	}
}

func TestExec(t *testing.T) {
	m := map[string]struct {
		R        Rover
		Expected string
	}{
		"LMLMLMLMM": {
			R: Rover{
				Heading{North},
				Point{1, 2},
				Point{5, 5},
			},
			Expected: "1 3 N",
		},
		"MMRMMRMRRM": {
			R: Rover{
				Heading{East},
				Point{3, 3},
				Point{5, 5},
			},
			Expected: "5 1 E",
		},
	}

	for c, v := range m {
		if err := v.R.Exec(c); err != nil {
			t.Errorf("unable to execute commands on rover: %s", err)
		}

		if v.R.String() != v.Expected {
			t.Errorf("rover %s is not at desired place %s", v.R, v.Expected)
		}
	}
}
