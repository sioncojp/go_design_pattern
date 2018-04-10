package adapter

import "fmt"

type FruitPrice interface {
	getPrice()
}

type Fruit struct {
	Name string
	Cost int
}

type FruitAdapter struct {
	Fruit
}

func (s *FruitAdapter) GetPrice() {
	s.Fruit.GetPrice()
}

func (s *Fruit) GetPrice() {
	fmt.Printf("%s: %d\n", s.Name, s.Cost)
}
