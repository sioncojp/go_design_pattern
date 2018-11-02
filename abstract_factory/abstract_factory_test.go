package abstract_factory

import (
	"testing"
	"fmt"
)

func TestAbstractFactory (t *testing.T) {
	wood := NewDoorFactory("wood")
	iron := NewDoorFactory("iron")

	fmt.Println(wood.makeDoor().getType())
	fmt.Println(wood.makeFittingExpert().getFitting())

	fmt.Println(iron.makeDoor().getType())
	fmt.Println(iron.makeFittingExpert().getFitting())
}
