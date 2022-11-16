package main

type ChristmasLightsOn struct {
	lightsConfiguration LightsConfiguration
}

func newLightsOn(lightsConfiguration LightsConfiguration) Light {
	return &ChristmasLightsOn{
		lightsConfiguration: lightsConfiguration,
	}
}

func (clo ChristmasLightsOn) Display() [][]string {
	lightsConfiguration := clo.lightsConfiguration
	lights := lightsConfiguration.lights
	for i := lightsConfiguration.rowX; i <= lightsConfiguration.rowY; i++ {
		lights[i] = changeLightRowOn(lightsConfiguration, lights[i])
	}
	return lights
}

func changeLightRowOn(configuration LightsConfiguration, row []string) []string {
	for i := configuration.columnX; i <= configuration.columnY; i++ {
		row[i] = TurnOn
	}
	return row
}
