package mars

// Heading represents the direction of rover
type Heading struct {
	Cursor int
}

// Right rotates 90˚ to the right of current direction
func (h *Heading) Right() {
	h.Cursor = (h.Cursor + 1) % len(Seq)
}

// Left rotates 90˚ to the left of current direction
func (h *Heading) Left() {
	n := h.Cursor - 1
	if n < 0 {
		n = len(Seq) - 1
	}
	h.Cursor = n
}
