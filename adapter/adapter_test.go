package adapter

import "testing"

func TestAdapter(t *testing.T) {
	t1 := &Taro{}
	t1.EnjoyWithClassmate()

	t2 := &TaroV2{*t1}
	t2.OrganizationClass()
}
