package kata_log

import (
	"reflect"
	"testing"
)

func TestDisplay(t *testing.T) {

	t.Run("Should display all lights on off", func(t *testing.T) {
		want := make([][]string, 1000)

		want = getLightsAllOff(want)
		got := Display(want, TurnOff, 0, 0, 0, 0)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Should display all lights on on", func(t *testing.T) {
		want := make([][]string, 1000)
		lights := make([][]string, 1000)

		want = getLightsAllOn(want)
		got := Display(getLightsAllOff(lights), TurnOn, 0, 0, 0, 0)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Should toggle the first line, turning off the ones that were on, "+
		"and turning on the ones the that were off", func(t *testing.T) {
		want := [][]string{
			{"0", "1", "0", "1"},
			{"1", "0", "1", "0"},
			{"0", "0", "1", "0"},
			{"1", "0", "1", "0"},
		}
		lights := [][]string{
			{"1", "1", "0", "1"},
			{"0", "0", "1", "0"},
			{"1", "0", "1", "0"},
			{"0", "0", "1", "0"},
		}

		got := Display(lights, Toggle, 0, 3, 0, 0)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func getLightsAllOn(want [][]string) [][]string {
	for i := range want {
		want[i] = getLightRowAllOn()
	}
	return want
}

func getLightRowAllOn() []string {
	lightRow := make([]string, 1000)
	for i := range lightRow {
		lightRow[i] = "1"
	}
	return lightRow
}

func getLightsAllOff(want [][]string) [][]string {
	for i := range want {
		want[i] = getLightRowAllOff()
	}
	return want
}

func getLightRowAllOff() []string {
	lightRow := make([]string, 1000)
	for i := range lightRow {
		lightRow[i] = "0"
	}
	return lightRow
}
