package adapter

import "testing"

func TestAdapterExtends(t *testing.T) {
	f1 := &Fruit{"Apple", 100}
	f2 := &FruitAdapter{*f1}

	f1.GetPrice()
	f2.GetPrice()
}
