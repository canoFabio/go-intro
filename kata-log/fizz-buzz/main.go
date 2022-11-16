package fizzBuzz

import (
	"os"
)

func main() {
	fizzBuzz := NewFizzBuzz(os.Stdout)
	fizzBuzz.FindFizzBuzz(1, 100)
}
