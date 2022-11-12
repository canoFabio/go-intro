package main

type ChristmasLightsToggle struct {
	lightsConfiguration LightsConfiguration
}

func newLightsToggle(lightsConfiguration LightsConfiguration) ChristmasLight {
	return &ChristmasLightsToggle{
		lightsConfiguration: lightsConfiguration,
	}
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
