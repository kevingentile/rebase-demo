package animal

// Human is a mamel with a name
type Human struct {
	Name string `json:"name"`
}

// Legs returns a humans leg count in our microverse.
func (h *Human) Legs() int {
	return 2
}

// NewHuman creates a human with name
func NewHuman(name string) *Human {
	return &Human{Name: name}
}

// Shark is a swimmer with a species
type Shark struct {
	Species string `json:"species"`
}

// Fins returns the count of fins of a shark in our microverse.
func (b *Shark) Fins() int {
	return 4
}

// NewShark creates a new fish with species
func NewShark(species string) *Shark {
	return &Shark{Species: species}
}
