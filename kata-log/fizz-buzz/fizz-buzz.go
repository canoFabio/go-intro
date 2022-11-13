package fizz_buzz

import (
	"fmt"
	"io"
)

type FizzBuzz struct {
	out io.Writer
}

func (fb *FizzBuzz) FindFizzBuzz(since, to int) {
	for i := since; i <= to; i++ {
		fmt.Fprintln(fb.out, i)
	}
}

func NewFizzBuzz(out io.Writer) *FizzBuzz {
	return &FizzBuzz{
		out: out,
	}
}
