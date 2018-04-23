package prototype

import "fmt"

type cloneabler interface {
	Get() string
	clone() cloneabler
}

type Person struct {
	cloneable cloneabler
}

type Pepper struct {
	Figure string
}

func Draw(s string) *Pepper {
	fmt.Printf("%sを書くよ\n", s)
	return &Pepper{Figure: s}
}

func Cut(s string) *Pepper {
	pepper := Draw(s)
	fmt.Printf("%sの形に切るお\n", pepper.Figure)
	return pepper
}

func (s *Person) Register(c cloneabler) {
	s.cloneable = c
}

func (s *Person) CreateClone() cloneabler {
	return s.cloneable.clone()
}

func (s *Pepper) Get() string {
	return s.Figure
}

// スライスならdeep copyをする必要がある
func (s *Pepper) clone() cloneabler {
	return &Pepper{s.Figure}
}
