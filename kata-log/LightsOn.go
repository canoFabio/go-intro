package kata_log

type ChristmasLightsOn struct {
	lightsConfiguration LightsConfiguration
}

func newLightsOn() ChristmasLight {
	return &ChristmasLightsOn{}
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
