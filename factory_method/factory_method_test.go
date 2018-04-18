package factory_method

import (
	"fmt"
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	users := []User{
		User{"Taro", "WGundam"},
		User{"Jiro", "GodGundam"},
	}

	for _, v := range users {
		s := NewGundam(v.Buy)
		fmt.Printf("%sは", v.Name)
		s.Create()
	}
}
