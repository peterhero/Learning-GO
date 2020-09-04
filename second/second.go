package main

import "fmt"

type car struct {
	gasPedal      uint16
	brakePedal    uint16
	steeringWheel int16
	topSpeedKmh   float64
}

func main() {
	car1 := car{gasPedal: 22341, // possible values only
		brakePedal:    0, // but this is better way
		steeringWheel: 12561,
		topSpeedKmh:   225.0}

	fmt.Println(car1.gasPedal)
}
