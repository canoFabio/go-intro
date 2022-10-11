package kata_log

type Configuration string

const (
	TurnOn  = "1"
	TurnOff = "0"
)

type LightsConfiguration struct {
	lights                       [][]string
	rowX, rowY, columnX, columnY int
}

type Light interface {
	Display() [][]string
}

type ChristmasLightsOn struct {
	lightsConfiguration LightsConfiguration
}
type ChristmasLightsOff struct {
	lightsConfiguration LightsConfiguration
}
type ChristmasLightsToggle struct {
	lightsConfiguration LightsConfiguration
}

func (clo ChristmasLightsOn) Display() [][]string {
	lightsConfiguration := clo.lightsConfiguration
	lights := lightsConfiguration.lights
	for i := range lights {
		lights[i] = changeAllLightsInARowState(lights[i], TurnOn)
	}
	return lights
}

func changeAllLightsInARowState(lightRow []string, lightState string) []string {
	for i := range lightRow {
		lightRow[i] = lightState
	}
	return lightRow
}

func (clo ChristmasLightsOff) Display() [][]string {
	lightsConfiguration := clo.lightsConfiguration
	return lightsConfiguration.lights
}

func (clo ChristmasLightsToggle) Display() [][]string {
	lightsConfiguration := clo.lightsConfiguration
	lights := lightsConfiguration.lights
	for i := lightsConfiguration.rowX; i <= lightsConfiguration.rowY; i++ {
		lights[i] = changeLightsRowInToggle(lightsConfiguration.columnX,
			lightsConfiguration.columnY,
			lights[i])
	}
	return lights
}

func changeLightsRowInToggle(columnX int, columnY int, light []string) []string {
	for j := columnX; j <= columnY; j++ {
		light[j] = toggleLight(light[j])
	}
	return light
}

func toggleLight(lightState string) string {
	if lightState == TurnOn {
		return TurnOff
	}
	return TurnOn
}
