package adapter

type ChairPerson interface {
	OrganizeClass()
}

type TaroV2 struct {
	Taro
}

type Taro struct{}

func (s *TaroV2) OrganizationClass() {
	println("クラスの面倒みるよ")
	s.Taro.EnjoyWithClassmate()
}

func (s *Taro) EnjoyWithClassmate() {
	println("わいわい")
}
