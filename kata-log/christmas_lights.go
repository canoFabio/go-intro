package kata_log

type Configuration string

const (
	TurnOn  = Configuration("1")
	TurnOff = Configuration("0")
	Toggle  = Configuration("2")
)

type LightsConfiguration struct {
	lights                       [][]string
	configuration                Configuration
	rowX, rowY, columnX, columnY int
}

type Lights interface {
	DisplayLightsConfiguration() [][]string
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

func (clo ChristmasLightsOn) DisplayLightsConfiguration() [][]string {
	lightsConfiguration := clo.lightsConfiguration
	lights := lightsConfiguration.lights
	for i := range lights {
		lights[i] = changeAllLightsInARowState(lights[i], lightsConfiguration.configuration.getValue())
	}
	return lights
}

func (clo ChristmasLightsOff) DisplayLightsConfiguration() [][]string {
	lightsConfiguration := clo.lightsConfiguration
	return lightsConfiguration.lights
}

func (clo ChristmasLightsToggle) DisplayLightsConfiguration() [][]string {
	lightsConfiguration := clo.lightsConfiguration
	lights := lightsConfiguration.lights
	for i := lightsConfiguration.rowX; i <= lightsConfiguration.rowY; i++ {
		lights[i] = changeLightsRowInToggle(lightsConfiguration.columnX, lightsConfiguration.columnY, lights[i])
	}
	return lights
}

func (c Configuration) getValue() string {
	return string(c)
}

func changeAllLightsInARowState(lightRow []string, lightState string) []string {
	for i := range lightRow {
		lightRow[i] = lightState
	}
	return lightRow
}

func changeLightsRowInToggle(columnX int, columnY int, light []string) []string {
	for j := columnX; j <= columnY; j++ {
		light[j] = toggleLight(light[j])
	}
	return light
}

func toggleLight(lightState string) string {
	if lightState == TurnOn.getValue() {
		return TurnOff.getValue()
	}
	return TurnOn.getValue()
}
