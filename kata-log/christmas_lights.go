package kata_log

type Configuration string

const (
	TurnOn  = Configuration("1")
	TurnOff = Configuration("0")
	Toggle  = Configuration("2")
)

func (c Configuration) getValue() string {
	return string(c)
}

func Display(lights [][]string, configuration Configuration, rowX, rowY, columnX, columnY int) [][]string {
	switch configuration {
	case TurnOn:
		return changeAllLightsForNewState(lights, configuration.getValue())
	case Toggle:
		return changeRangeOfLights(lights, rowX, rowY, columnX, columnY)
	default:
		return lights
	}
}

func changeAllLightsForNewState(lights [][]string, lightState string) [][]string {
	for i := range lights {
		lights[i] = changeAllLightsInARowState(lights[i], lightState)
	}
	return lights
}

func changeAllLightsInARowState(lightRow []string, lightState string) []string {
	for i := range lightRow {
		lightRow[i] = lightState
	}
	return lightRow
}

func changeRangeOfLights(lights [][]string, rowX, rowY, columnX, columnY int) [][]string {
	for i := rowX; i <= rowY; i++ {
		lights[i] = changeLightsRowInToggle(columnX, columnY, lights[i])
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
	if lightState == TurnOn.getValue() {
		return TurnOff.getValue()
	}
	return TurnOn.getValue()
}
