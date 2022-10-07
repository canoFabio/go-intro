package kata_log

import (
	"reflect"
	"testing"
)

func TestDisplay(t *testing.T) {

	t.Run("Should display all lights on off", func(t *testing.T) {
		want := make([][]string, 1000)

		want = getLightsAllOff(want)
		got := Display(want, false)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Should display all lights on on", func(t *testing.T) {
		want := make([][]string, 1000)
		lights := make([][]string, 1000)

		want = getLightsAllOn(want)
		got := Display(getLightsAllOff(lights), true)

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
