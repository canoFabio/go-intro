package kata_log

func Display(lights [][]string, turnOn bool) [][]string {
	if turnOn {
		return changeAllLightsForNewState(lights, "1")
	}
	return lights
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
