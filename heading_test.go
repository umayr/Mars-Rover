package mars

import "testing"

func TestRight(t *testing.T) {
	h := Heading{}

	if Seq[h.Cursor] != "N" {
		t.Error("Heading should be initialised with North")
	}

	h.Right()
	if Seq[h.Cursor] != "E" {
		t.Error("Heading should be East when rotated right from North")
	}

	h.Right()
	if Seq[h.Cursor] != "S" {
		t.Error("Heading should be South when rotated right from East")
	}

	h.Right()
	if Seq[h.Cursor] != "W" {
		t.Error("Heading should be West when rotated right from South")
	}

	h.Right()
	if Seq[h.Cursor] != "N" {
		t.Error("Heading should be North when rotated right from West")
	}
}

func TestLeft(t *testing.T) {
	h := Heading{}

	if Seq[h.Cursor] != "N" {
		t.Error("Heading should be initialised with North")
	}

	h.Left()
	if Seq[h.Cursor] != "W" {
		t.Error("Heading should be West when rotated right from North")
	}

	h.Left()
	if Seq[h.Cursor] != "S" {
		t.Error("Heading should be South when rotated right from West")
	}

	h.Left()
	if Seq[h.Cursor] != "E" {
		t.Error("Heading should be East when rotated right from South")
	}

	h.Left()
	if Seq[h.Cursor] != "N" {
		t.Error("Heading should be North when rotated right from East")
	}
}
