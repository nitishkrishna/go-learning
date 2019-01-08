package main

import (
	"fmt"
	// Subpackage
	"math/rand"
	"time"
)

// Vid 3 Types

func add(x, y float32) float32 {
	return x + y
}

func mult(first, second string) (string, string) {
	return first, second
}

// vid 6 Structs

type car struct {
	gasPedal      uint16 // 0-65k
	brakePedal    uint16
	steeringWheel int16 // -32k to 32k
	topSpeedKmph  float64
}

// Vid 7 Methods

const usixteenbitmax float64 = 65535
const mphToKmph float64 = 1.60934

// Methods for Car Struct
func (c car) kmph() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmph / usixteenbitmax)
}
func (c car) mph() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmph / usixteenbitmax / mphToKmph)
}

// Pointer receiver funcs
func (c *car) newTopSpeed(newSpeed float64) {
	c.topSpeedKmph = newSpeed
}

// Non Struct Function - Pass Struct instance and overwrite
func changeTopSpeed(c car, newSpeed float64) car {
	c.topSpeedKmph = newSpeed
	return c
}

func main() {

	// Vid 2 Syntax
	// Seed the random val generator
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Random Int bw 1 and 100", rand.Intn(100))

	// Vid 3 Types
	// Define and add vars - Typing Info
	a, b := 5.5, 6.6
	// Cast to float32 for func to accept
	fmt.Println(add(float32(a), float32(b)))

	// Func returns multiple vals
	w1, w2 := "Hello", "World"
	fmt.Println(mult(w1, w2))

	// Vid 4 Pointers

	x := 15
	m := &x // Memory Add of x

	fmt.Println("Mem of x:", m)
	fmt.Println("Val of x:", *m)

	*m = 5

	fmt.Println("Val of x now:", x)
	// Square using pointer
	*m = *m * *m
	fmt.Println("Val of x squared:", x)
	fmt.Println("Val of a now:", m)

	// vid 6 Structs
	// Structs and Classes are same
	// Declare instance
	firstCar := car{
		gasPedal:      22341,
		brakePedal:    0,
		steeringWheel: 10090,
		topSpeedKmph:  225.0}

	// Print attribute
	fmt.Println("Gas Pedal:", firstCar.gasPedal)

	// Vid 7 Methods for Structs

	// Two types of Methods -
	// Value receivers : just receive values and perform calcs

	fmt.Println("KMPH Speed:", firstCar.kmph())
	fmt.Println("MPH Speed:", firstCar.mph())

	// Pointer receivers: to modify values received
	// Vid 8
	firstCar.newTopSpeed(500)
	fmt.Println("New KMPH Speed:", firstCar.kmph())
	fmt.Println("New MPH Speed:", firstCar.mph())

	firstCar = changeTopSpeed(firstCar, 600)
	fmt.Println("New KMPH Speed:", firstCar.kmph())
	fmt.Println("New MPH Speed:", firstCar.mph())

}
