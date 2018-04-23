package prototype

import "testing"

func TestPrototype(t *testing.T) {
	figure := "雪の結晶"
	pepper := Cut(figure)

	you := &Person{}
	you.Register(pepper)

	cloned := you.CreateClone()

	if cloned.Get() != pepper.Get() {
		t.Errorf("failed")
	}
}
