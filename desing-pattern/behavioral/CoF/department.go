package CoF

type Department interface {
	execute(*Patient)
	setNext(Department)
}
