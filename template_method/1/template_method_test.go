package template_method

import "testing"

func TestLion(t *testing.T) {
	s := &Lion{}
	Behavior(s)
}

func TestBird(t *testing.T) {
	s := &Bird{}
	Behavior(s)
}
