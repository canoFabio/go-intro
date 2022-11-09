package kata_log

type ChristmasLightsOff struct {
	lightsConfiguration LightsConfiguration
}

func newLightsOff(lightsConfiguration LightsConfiguration) ChristmasLight {
	return &ChristmasLightsOff{
		lightsConfiguration: lightsConfiguration,
	}
}

func (clo ChristmasLightsOff) Display() [][]string {
	lightsConfiguration := clo.lightsConfiguration
	return lightsConfiguration.lights
}
