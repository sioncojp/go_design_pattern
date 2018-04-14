package template_method

// 小文字にしてプライベートにする
type animal interface {
	eat()
	sleep()
}

type Animal struct {
	animal animal
}

func (s *Animal) Behavior() {
	s.animal.eat()
	s.animal.sleep()
}

type Lion struct{}

func (s *Lion) eat() {
	println("肉を食べる")
}

func (s *Lion) sleep() {
	println("草原で寝る")
}

type Bird struct{}

func (s *Bird) eat() {
	println("虫を食べる")
}

func (s *Bird) sleep() {
	println("巣で寝る")
}
