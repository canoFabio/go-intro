package kata_log

import (
	"reflect"
	"testing"
)

func TestDisplay(t *testing.T) {
	want := make([][]string, 1000)

	want = getLightsAllOff(want)
	got := Display(want)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func getLightsAllOff(want [][]string) [][]string {
	for range want {
		want = append(want, getLightRowAllOff())
	}
	return want
}

func getLightRowAllOff() []string {
	lightRow := make([]string, 1000)
	for range lightRow {
		lightRow = append(lightRow, "0")
	}
	return lightRow
}
