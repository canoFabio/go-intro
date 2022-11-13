package fizz_buzz

import (
	"fmt"
	"io"
	"strconv"
)

type FizzBuzz struct {
	out io.Writer
}

func NewFizzBuzz(out io.Writer) *FizzBuzz {
	return &FizzBuzz{
		out: out,
	}
}

func (fb *FizzBuzz) FindFizzBuzz(since, to int) {
	for i := since; i <= to; i++ {
		fmt.Fprintln(fb.out, fb.findValueToWrite(i))
	}
}

func (fb *FizzBuzz) findValueToWrite(number int) string {
	if fb.isMultiplyOfThree(number) {
		return "Fizz"
	}
	if fb.isMultiplyOfFive(number) {
		return "Buzz"
	}
	return strconv.Itoa(number)
}

func (fb *FizzBuzz) isMultiplyOfFive(number int) bool {
	return number%5 == 0
}

func (fb *FizzBuzz) isMultiplyOfThree(number int) bool {
	return number%3 == 0
}
