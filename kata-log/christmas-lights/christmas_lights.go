package main

const (
	TurnOn  = "1"
	TurnOff = "0"
)

type LightsConfiguration struct {
	lights                       [][]string
	rowX, rowY, columnX, columnY int
}
