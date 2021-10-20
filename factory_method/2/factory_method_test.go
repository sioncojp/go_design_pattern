package factory_method

import "testing"

func TestFactoryMethod(t *testing.T) {
	guns := []iGun{
		getGun("ak47"),
		getGun("musket"),
	}

	for _, v := range guns {
		printDetails(v)
		printDetails(v)
	}
}
