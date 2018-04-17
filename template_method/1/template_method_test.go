package template_method

import "testing"

func TestTemplateMethod1(t *testing.T) {
	s := &Taro{}
	s.Create(s)
}
