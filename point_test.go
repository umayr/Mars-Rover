package mars

import (
	"testing"
)

func TestOut(t *testing.T) {
	s := &Point{3, 5}
	m := map[Point]bool{
		Point{0, 0}: true,
		Point{0, 1}: true,
		Point{0, 2}: true,
		Point{0, 3}: true,
		Point{0, 4}: true,
		Point{0, 5}: true,
		Point{1, 0}: true,
		Point{1, 1}: true,
		Point{1, 2}: true,
		Point{1, 3}: true,
		Point{1, 4}: true,
		Point{1, 5}: true,
		Point{2, 0}: true,
		Point{2, 1}: true,
		Point{2, 2}: true,
		Point{2, 3}: true,
		Point{2, 4}: true,
		Point{2, 5}: true,
		Point{3, 0}: true,
		Point{3, 1}: true,
		Point{3, 2}: true,
		Point{3, 3}: true,
		Point{3, 4}: true,
		Point{4, 1}: true,
		Point{4, 2}: true,
		Point{4, 3}: true,
		Point{4, 4}: true,
		Point{5, 1}: true,
		Point{5, 2}: true,
		Point{5, 3}: true,
		Point{5, 4}: true,
		Point{6, 0}: true,
		Point{0, 6}: true,
		Point{3, 5}: false,
		Point{3, 6}: false,
		Point{3, 7}: false,
		Point{3, 8}: false,
		Point{3, 9}: false,
		Point{4, 5}: false,
		Point{4, 6}: false,
		Point{4, 7}: false,
		Point{4, 8}: false,
		Point{4, 9}: false,
		Point{5, 5}: false,
		Point{5, 6}: false,
		Point{5, 7}: false,
		Point{5, 8}: false,
		Point{5, 9}: false,
	}

	for p, v := range m {
		if r := s.Out(&p); r != v {
			if r {
				t.Errorf("point %s is not out of bound of point %s", s, p)
			} else {
				t.Errorf("point %s is out of bound of point %s", s, p)
			}
		}
	}
}
