package kata_log

import (
	"reflect"
	"testing"
)

func TestDisplay(t *testing.T) {

	lightsTest := []struct {
		name      string
		light     ChristmasLight
		hasLights [][]string
	}{
		{name: "Should display all lights on off", light: ChristmasLightsOff{lightsConfiguration: getLightsConfigAllOff()}, hasLights: getLightsAllOff()},
		{name: "Should display all lights on on", light: ChristmasLightsOn{lightsConfiguration: getLightsConfigAllOff()}, hasLights: getLightsAllOn()},
		{name: "Should toggle the first line, turning off the ones that were on, " +
			"and turning on the ones the that were off", light: ChristmasLightsToggle{lightsConfiguration: getSomeRandomConfigLights()}, hasLights: getResultForSomeRandomLights()},
	}

	for _, lt := range lightsTest {
		t.Run(lt.name, func(t *testing.T) {
			got := lt.light.Display()
			if !reflect.DeepEqual(got, lt.hasLights) {
				t.Errorf("got %v want %v", got, lt.hasLights)
			}
		})
	}
}

func getResultForSomeRandomLights() [][]string {
	return [][]string{
		{"0", "1", "0", "1"},
		{"1", "0", "1", "0"},
		{"0", "0", "1", "0"},
		{"1", "0", "1", "0"},
	}
}

func getSomeRandomConfigLights() LightsConfiguration {
	return LightsConfiguration{
		lights:  getSomeRandomLights(),
		rowX:    0,
		rowY:    3,
		columnX: 0,
		columnY: 0}
}

func getSomeRandomLights() [][]string {
	return [][]string{
		{"1", "1", "0", "1"},
		{"0", "0", "1", "0"},
		{"1", "0", "1", "0"},
		{"0", "0", "1", "0"},
	}
}

func getLightsConfigAllOff() LightsConfiguration {
	return LightsConfiguration{
		lights:  getLightsAllOff(),
		rowX:    0,
		rowY:    0,
		columnX: 0,
		columnY: 0}
}

func getLightsAllOn() [][]string {
	want := make([][]string, 1000)
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

func getLightsAllOff() [][]string {
	want := make([][]string, 1000)
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
