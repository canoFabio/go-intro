package kata_log

import "fmt"

const (
	lightsOff    = "lights_off"
	lightsOn     = "lights_on"
	lightsToggle = "lights_toggle"
)

func getChristmasLights(christmasLightsType string, lightsConfiguration LightsConfiguration) (ChristmasLight, error) {
	if christmasLightsType == lightsOn {
		return newLightsOn(lightsConfiguration), nil
	}

	if christmasLightsType == lightsOff {
		return newLightsOff(lightsConfiguration), nil
	}

	if christmasLightsType == lightsToggle {
		return newLightsToggle(lightsConfiguration), nil
	}

	return nil, fmt.Errorf("wrong gun type passed")
}
