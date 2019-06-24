package main

import "fmt"

const (
	_ = iota
	TrafficLightStateRedLight = "red"
	TrafficLightStateGreenLight = "green"
	TrafficLightStateYellowLight = "yellow"
)

func main() {
	fmt.Println("Red Light State Code:", TrafficLightStateRedLight)
	fmt.Println("Green Light State Code:", TrafficLightStateGreenLight)
	fmt.Println("Yellow Light State Code:", TrafficLightStateYellowLight)
}