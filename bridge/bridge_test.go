package bridge

import (
	"testing"
)

func TestBridge(t *testing.T) {
	sword := &Sword{}
	bow := &Bow{}

	s := &Smith{sword}
	w := &Warrior{bow}

	cases := []struct {
		value    string
		expected string
	}{
		{
			Do(s),
			"sword Repair",
		},
		{
			Do(w),
			"bow Attack",
		},
	}
	for _, c := range cases {
		if c.value != c.expected {
			t.Errorf("want %s got %s", c.expected, c.value)
		}
	}
}
