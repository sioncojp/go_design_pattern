package bridge

type Weapon interface {
	Attack() string
	Repair() string
}

type Bow struct{}

func (b *Bow) Attack() string {
	return "bow Attack"
}
func (b *Bow) Repair() string {
	return "bow Repair"
}

type Sword struct{}

func (s *Sword) Attack() string {
	return "sword Attack"
}
func (s *Sword) Repair() string {
	return "sword Repair"
}

type WeaponHandler interface {
	Handle() string
}

type Warrior struct {
	weapon Weapon
}

func (w *Warrior) Handle() string {
	return w.weapon.Attack()
}

type Smith struct {
	weapon Weapon
}

func (s *Smith) Handle() string {
	return s.weapon.Repair()
}

func Do(w WeaponHandler) string {
	return w.Handle()
}
