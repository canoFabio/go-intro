package main

import "fmt"

func main() {
	santaClausChristmasLights := make(map[*LightsConfiguration]string)

	santaClausChristmasLights[&LightsConfiguration{rowX: 887, columnX: 9, rowY: 959, columnY: 629}] = lightsOn
	santaClausChristmasLights[&LightsConfiguration{rowX: 454, columnX: 398, rowY: 844, columnY: 448}] = lightsOn
	santaClausChristmasLights[&LightsConfiguration{rowX: 539, columnX: 243, rowY: 559, columnY: 965}] = lightsOff
	santaClausChristmasLights[&LightsConfiguration{rowX: 370, columnX: 819, rowY: 676, columnY: 868}] = lightsOff
	santaClausChristmasLights[&LightsConfiguration{rowX: 154, columnX: 40, rowY: 370, columnY: 997}] = lightsOff
	santaClausChristmasLights[&LightsConfiguration{rowX: 301, columnX: 3, rowY: 808, columnY: 453}] = lightsOff
	santaClausChristmasLights[&LightsConfiguration{rowX: 351, columnX: 678, rowY: 951, columnY: 908}] = lightsOn
	santaClausChristmasLights[&LightsConfiguration{rowX: 720, columnX: 196, rowY: 897, columnY: 994}] = lightsToggle
	santaClausChristmasLights[&LightsConfiguration{rowX: 831, columnX: 394, rowY: 904, columnY: 860}] = lightsToggle

	result := DisplaySantaClausLights(santaClausChristmasLights)

	fmt.Printf("Cant of lighst on %d", countLightsOn(result))
	fmt.Printf("%v", result)

}

func countLightsOn(lights [][]string) int {
	var cant int
	for i := range lights {
		row := lights[i]
		for g := range row {
			if row[g] == TurnOn {
				cant++
			}
		}
	}
	return cant
}

func DisplaySantaClausLights(lights map[*LightsConfiguration]string) [][]string {
	light := getLightsAllOff2()
	for key, value := range lights {
		key.lights = light
		christmasLights, _ := getChristmasLights(value, *key)
		light = christmasLights.Display()
	}

	return light
}

func getLightsAllOff2() [][]string {
	want := make([][]string, 1000)
	for i := range want {
		want[i] = getLightRowAllOff2()
	}
	return want
}

func getLightRowAllOff2() []string {
	lightRow := make([]string, 1000)
	for i := range lightRow {
		lightRow[i] = "0"
	}
	return lightRow
}
