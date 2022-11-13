package main

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

	switch {
	case fb.isMultiplyOfThree(number) && fb.isMultiplyOfFive(number):
		return "FizzBuzz"
	case fb.isMultiplyOfThree(number):
		return "Fizz"
	case fb.isMultiplyOfFive(number):
		return "Buzz"
	default:
		return strconv.Itoa(number)
	}

}

func (fb *FizzBuzz) isMultiplyOfFive(number int) bool {
	return number%5 == 0
}

func (fb *FizzBuzz) isMultiplyOfThree(number int) bool {
	return number%3 == 0
}
