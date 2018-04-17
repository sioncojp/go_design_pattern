package template_method

// 小文字にしてプライベートにする
type kit interface {
	cut()
	build()
	paint()
}

type Gundam struct {
}

func (s *Gundam) Create(kit kit) {
	kit.cut()
	kit.build()
	kit.paint()
}

type Taro struct {
	*Gundam
}

func (s *Taro) cut() {
	println("ニッパーで切る")
}

func (s *Taro) build() {
	println("組み立てる")
}

func (s *Taro) paint() {
	println("色を塗る")
}
