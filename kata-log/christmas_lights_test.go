package kata_log

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhenLightsConfigurationIsAllOff_ThenAllLightsShouldBeZero(t *testing.T) {
	lightConfiguration := getLightsConfigAllOff()
	want := getLightsAllOff()
	christmasLights, _ := getChristmasLights(lightsOff, lightConfiguration)

	got := christmasLights.Display()

	assert.Equal(t, got, want)

}

func TestWhenLightsConfigurationIsAllOn_ThenAllLightsShouldBeOne(t *testing.T) {
	lightConfiguration := getLightsConfigAllOff()
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
