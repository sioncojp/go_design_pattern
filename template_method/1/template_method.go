package template_method

// 小文字にしてプライベートにする
type animal interface {
	eat()
	sleep()
}

type Animal struct{}

func Behavior(a animal) {
	a.eat()
	a.sleep()
}

type Lion struct {
	*Animal
}

func (s *Lion) eat() {
	println("肉を食べる")
}

func (s *Lion) sleep() {
	println("草原で寝る")
}

type Bird struct {
	*Animal
}

func (s *Bird) eat() {
	println("虫を食べる")
}

func (s *Bird) sleep() {
	println("巣で寝る")
}
