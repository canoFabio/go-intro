package simple_factory

type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}
