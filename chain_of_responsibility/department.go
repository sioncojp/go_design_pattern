package chain_ob_responsibility

type department interface {
	execute(*patient)
	setNext(department)
}
