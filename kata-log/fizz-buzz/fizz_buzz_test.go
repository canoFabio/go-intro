package fizz_buzz

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhenFizzBuzz_ThenPrintTheNumberOne(t *testing.T) {
	//given
	since := 1
	to := 1
	buffer := &bytes.Buffer{}
	underTest := NewFizzBuzz(buffer)
	want := `1
`
	//when
	underTest.FindFizzBuzz(since, to)
	got := buffer.String()
	//then
	assert.Equal(t, want, got)
}

func TestWhenFizzBuzz_ThenPrintSinceTheNumberOneToTheNumberTwo(t *testing.T) {
	//given
	since := 1
	to := 2
	buffer := &bytes.Buffer{}
	underTest := NewFizzBuzz(buffer)
	want := `1
2
`
	//when
	underTest.FindFizzBuzz(since, to)
	got := buffer.String()
	//then
	assert.Equal(t, want, got)
}

func TestWhenFizzBuzzWithRangeFromOneToThree_ThenPrintSinceTheNumberOneToTheNumberTwoAndPrintFizzInsteadThree(t *testing.T) {
	//given
	since := 1
	to := 3
	buffer := &bytes.Buffer{}
	underTest := NewFizzBuzz(buffer)
	want := `1
2
Fizz
`
	//when
	underTest.FindFizzBuzz(since, to)
	got := buffer.String()
	//then
	assert.Equal(t, want, got)
}

func TestWhenFizzBuzzWithRangeFromOneToFive_ThenPrintSinceTheNumberOneToTheNumberTwoAndPrintFizzInsteadThreeAndPrintBuzzInsteadFive(t *testing.T) {
	//given
	since := 1
	to := 5
	buffer := &bytes.Buffer{}
	underTest := NewFizzBuzz(buffer)
	want := `1
2
Fizz
4
Buzz
`
	//when
	underTest.FindFizzBuzz(since, to)
	got := buffer.String()
	//then
	assert.Equal(t, want, got)
}
