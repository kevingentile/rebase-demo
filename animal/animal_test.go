package animal

import "testing"

func TestHuman(t *testing.T) {
	h := NewHuman("Bob")

	if h.Legs() != 2 {
		t.Errorf("expected 2 legs, got %d", h.Legs())
	}

	if h.Name != "Bob" {
		t.Errorf("expected name \"Bob\", got \"%s\"", h.Name)
	}
}

func TestShark(t *testing.T) {
	s := NewShark("hammerhead")
	if s.Fins() != 4 {
		t.Errorf("expected 4 fins, got %d", s.Fins())
	}

	if s.Species != "hammerhead" {
		t.Errorf("expected species \"hammerhead\", got \"%s\"", s.Species)
	}
}
