package template_method

import "testing"

func TestTemplateMethod2(t *testing.T) {
	s := &Gundam{new(Taro)}
	s.Create()
}
