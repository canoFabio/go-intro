package kata_log

func Display(lights [][]string, turnOn bool) [][]string {
	if turnOn {
		for i := range lights {
			lightRow := lights[i]
			for f := range lightRow {
				lightRow[f] = "1"
			}
			lights[i] = lightRow
		}
		return lights
	}
	return lights
}
