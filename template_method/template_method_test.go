package template_method

import "testing"

func TestLion(t *testing.T) {
	s := &Lion{}
	s.Behavior(s)
}

func TestBird(t *testing.T) {
	s := &Bird{}
	s.Behavior(s)
}
