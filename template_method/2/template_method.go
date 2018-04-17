package template_method

// 小文字にしてプライベートにする
type kit interface {
	cut()
	build()
	paint()
}

type Gundam struct {
	kit
}

func (s *Gundam) Create() {
	s.kit.cut()
	s.kit.build()
	s.kit.paint()
}

type Taro struct{}

func (s *Taro) cut() {
	println("ニッパーで切る")
}

func (s *Taro) build() {
	println("組み立てる")
}

func (s *Taro) paint() {
	println("色を塗る")
}
