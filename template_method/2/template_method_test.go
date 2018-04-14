package template_method

import "testing"

func TestLion(t *testing.T) {
	s := &Animal{new(Lion)}
	s.Behavior()
}

func TestBird(t *testing.T) {
	s := &Animal{new(Bird)}
	s.Behavior()
}
