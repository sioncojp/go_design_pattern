package abstract_factory

type AbstractFactroy interface {
	makeDoor() Door
	makeFittingExpert() Door
}

type Door interface {
	getType() string
	getFitting() string
}

type Wood struct{
	Type string
	Fitting string
}
type Iron struct{
	Type string
	Fitting string
}

func (w *Wood) makeDoor() Door {
	w.Type = "私は木のドアです"
	return w
}

func (i *Iron) makeDoor() Door {
	i.Type = "私は鉄のドアです"
	return i
}

func (w *Wood) makeFittingExpert() Door {
	w.Fitting = "私は木のドアが取り付けれます"
	return w
}

func (i *Iron) makeFittingExpert() Door {
	i.Fitting = "私は鉄のドアが取り付けれます"
	return i
}

func (w *Wood) getType() string {
	return w.Type
}

func (i *Iron) getType() string {
	return i.Type
}

func (w *Wood) getFitting() string {
	return w.Fitting
}

func (i *Iron) getFitting() string {
	return i.Fitting
}

func NewDoorFactory(s string) AbstractFactroy {
	switch s {
	case "wood":
		return &Wood{}
	case "iron":
		return &Iron{}
	}
	panic("you shoud choice wood or iron")
}