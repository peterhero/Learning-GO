package main

import ("fmt"
		"math"
		"math/rand"
		"time"
		)
func add(a, b float64) float64 {
	return a+b
}

func main()  {
	fmt.Println("Hello GO!")
	fmt.Println("The square root of 16 is", math.Sqrt(16))
	rand.Seed(time.Now().UnixNano())
	fmt.Println("I have rolled the dice, there is a ", rand.Intn(6) + 1, "on it")

	var num1 float64 = 3.1
	var num2 float64 = 6.8
	num3, num4 := 3.1, 6.8
	fmt.Println(add(num1, num2))
	fmt.Println(add(num3, num4))
}