package christmas_lights

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhenLightsConfigurationIsAllOff_ThenAllLightsShouldBeZero(t *testing.T) {
	lightConfiguration := getLightsConfigAllLights()
	want := getLightsAllOff()
	christmasLights, _ := getChristmasLights(lightsOff, lightConfiguration)

	got := christmasLights.Display()

	assert.Equal(t, got, want)

}

func TestWhenLightsConfigurationIsAllOn_ThenAllLightsShouldBeOne(t *testing.T) {
	lightConfiguration := getLightsConfigAllLights()
	want := getLightsAllOn()
	christmasLights, _ := getChristmasLights(lightsOn, lightConfiguration)

	got := christmasLights.Display()

	assert.Equal(t, got, want)

}

func TestWhenLightsConfigurationIsToggleTheFirstLine_ThenTheOnesShouldBeZeroAndTheZerosShouldBeOnes(t *testing.T) {
	lightConfiguration := getSomeRandomConfigLights()
	want := getResultForSomeRandomLights()
	christmasLights, _ := getChristmasLights(lightsToggle, lightConfiguration)

	got := christmasLights.Display()

	assert.Equal(t, got, want)

}

func TestWhenLightsConfigurationTurnOffTheMiddleLights_ThenTheMiddleLightsShouldBeZeros(t *testing.T) {
	lightConfiguration := getConfigLightsForTheMiddleLightsOff()
	want := getResultForTheMiddleLightsOff()
	christmasLights, _ := getChristmasLights(lightsOff, lightConfiguration)

	got := christmasLights.Display()

	assert.Equal(t, got, want)
	t.Logf("got %v want %v", got, want)
}

func TestWhenLightsConfigurationTurnOnTheMiddleLights_ThenTheMiddleLightsShouldBeZeros(t *testing.T) {
	lightConfiguration := getConfigLightsForTheMiddleLightsOn()
	want := getResultForTheMiddleLightsOn()
	christmasLights, _ := getChristmasLights(lightsOn, lightConfiguration)

	got := christmasLights.Display()

	assert.Equal(t, want, got)
	t.Logf("got %v want %v", got, want)
}

func getResultForTheMiddleLightsOn() [][]string {
	return [][]string{
		{"0", "0", "0", "0"},
		{"0", "1", "1", "0"},
		{"0", "1", "1", "0"},
		{"0", "0", "0", "0"},
	}
}

func getConfigLightsForTheMiddleLightsOn() LightsConfiguration {
	return LightsConfiguration{
		lights:  get4x4LightsAllOff(),
		rowX:    1,
		rowY:    2,
		columnX: 1,
		columnY: 2}
}

func get4x4LightsAllOff() [][]string {
	return [][]string{
		{"0", "0", "0", "0"},
		{"0", "0", "0", "0"},
		{"0", "0", "0", "0"},
		{"0", "0", "0", "0"},
	}
}

func getConfigLightsForTheMiddleLightsOff() LightsConfiguration {
	return LightsConfiguration{
		lights:  get4x4LightsAllOn(),
		rowX:    1,
		rowY:    2,
		columnX: 1,
		columnY: 2}

}

func get4x4LightsAllOn() [][]string {
	return [][]string{
		{"1", "1", "1", "1"},
		{"1", "1", "1", "1"},
		{"1", "1", "1", "1"},
		{"1", "1", "1", "1"},
	}
}

func getResultForTheMiddleLightsOff() [][]string {
	return [][]string{
		{"1", "1", "1", "1"},
		{"1", "0", "0", "1"},
		{"1", "0", "0", "1"},
		{"1", "1", "1", "1"},
	}
}

func getLightsConfigAllLights() LightsConfiguration {
	return LightsConfiguration{
		lights:  getLightsAllOff(),
		rowX:    0,
		rowY:    999,
		columnX: 0,
		columnY: 999}
}

func getLightsAllOff() [][]string {
	want := make([][]string, 1000)
	for i := range want {
		want[i] = getLightRowAllOff()
	}
	return want
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

func getLightRowAllOff() []string {
	lightRow := make([]string, 1000)
	for i := range lightRow {
		lightRow[i] = "0"
	}
	return lightRow
}
