// Package jedlik organizes races between various types of remote controlled cars.
package jedlik

import "strconv"

// TODO: define the 'Drive()' method
// Drive updates the number of meters driven based on the car's speed,
// and reduces the battery according to the battery drainage:
func (car *Car) Drive() {
	if car.battery >= car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
}

// TODO: define the 'DisplayDistance() string' method
// DisplayDistance returns the distance as displayed on the LED display as a string.
func (car Car) DisplayDistance() string {
	return "Driven " + strconv.Itoa(car.distance) + " meters"
}

// TODO: define the 'DisplayBattery() string' method
// DisplayBattery returns the battery percentage as displayed on the LED display as a string.
func (car Car) DisplayBattery() string {
	return "Battery at " + strconv.Itoa(car.battery) + "%"
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
// CanFinish returns true if the car can finish the race; otherwise, returns false.
func (car Car) CanFinish(trackDistance int) bool {
	if float64(car.battery)-(float64(trackDistance)/float64(car.speed))*float64(car.batteryDrain) < 0.0 {
		return false
	}
	return true
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
