package main

import "fmt"

const u16max = 65535
const kmhMultiple float64 = 1.60934

type car struct {
	gasPedal      uint16
	brakePedal    uint16
	steeringWheel int16
	topSpeedKmh   float64
}

func (c car) kmh() float64{
	return float64(c.gasPedal) * (c.topSpeedKmh/float64(u16max))
}

func (c car) mph() float64{
	return float64(c.gasPedal) * (c.topSpeedKmh/float64(u16max))/kmhMultiple
}

func (c *car) newTopSpeed(newSpeed float64) {
	c.topSpeedKmh = newSpeed
}

func main() {
	car1 := car{gasPedal: 65000, // possible values only
		brakePedal:    0, // but this is better way
		steeringWheel: 12561,
		topSpeedKmh:   225.0}

	fmt.Println(car1.gasPedal)
	fmt.Println(car1.kmh())
	fmt.Println(car1.mph())

	car1.newTopSpeed(500)
	fmt.Println(car1.kmh())
	fmt.Println(car1.mph())
}