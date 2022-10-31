package kata_log

type ChristmasLightsOff struct {
	lightsConfiguration LightsConfiguration
}

func newLightsOff() ChristmasLight {
	return &ChristmasLightsOff{}
}

func (clo ChristmasLightsOff) Display() [][]string {
	lightsConfiguration := clo.lightsConfiguration
	return lightsConfiguration.lights
}
