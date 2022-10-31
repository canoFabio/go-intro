package kata_log

import "fmt"

func getChristmasLights(christmasLightsType string) (ChristmasLight, error) {
	if christmasLightsType == "lights_on" {
		return newLightsOn(), nil
	}

	if christmasLightsType == "lights_off" {
		return newLightsOff(), nil
	}

	if christmasLightsType == "lights_toggle" {
		return newLightsToggle(), nil
	}

	return nil, fmt.Errorf("wrong gun type passed")
}
