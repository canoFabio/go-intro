package main

type ChristmasLightsOff struct {
	lightsConfiguration LightsConfiguration
}

func newLightsOff(lightsConfiguration LightsConfiguration) Light {
	return &ChristmasLightsOff{
		lightsConfiguration: lightsConfiguration,
	}
}

func (clo ChristmasLightsOff) Display() [][]string {
	lightsConfiguration := clo.lightsConfiguration
	lights := lightsConfiguration.lights
	for i := lightsConfiguration.rowX; i <= lightsConfiguration.rowY; i++ {
		lights[i] = changeLightRowOff(lightsConfiguration, lights[i])
	}
	return lights
}

func changeLightRowOff(configuration LightsConfiguration, row []string) []string {
	for i := configuration.columnX; i <= configuration.columnY; i++ {
		row[i] = TurnOff
	}
	return row
}
