// Package thefarm calculates the amount of fodder that each cow should get
// and implements a neccesary error handling.
package thefarm

import (
	"errors"
	"strconv"
)

// TODO: define the 'DivideFood' function
// DivideFood returns the amount of food per cow as a float64 or an error if one occurred.
func DivideFood(fc FodderCalculator, nCows int) (float64, error) {
	amount, errAmount := fc.FodderAmount(nCows)
	if errAmount != nil {
		return 0.0, errAmount
	}
	factor, errFactor := fc.FatteningFactor()
	if errFactor != nil {
		return 0.0, errFactor
	}
	return float64((amount * factor) / float64(nCows)), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
// ValideInputAndDivideFood checks the number of cows.
func ValidateInputAndDivideFood(fc FodderCalculator, nCows int) (float64, error) {
	if nCows <= 0 {
		return 0.0, errors.New("invalid number of cows")
	}
	value, err := DivideFood(fc, nCows)
	return value, err
}

type InvalidCowsError struct {
	nCows   int
	message string
}

// Error returns a custom message.
func (i *InvalidCowsError) Error() string {
	return strconv.Itoa(i.nCows) + " cows are invalid: " + i.message
}

// TODO: define the 'ValidateNumberOfCows' function
// ValidateNumberOfCows accepts the number of cows as an integer and returns an error (or nil).
func ValidateNumberOfCows(nCows int) error {
	if nCows < 0 {
		return &InvalidCowsError{nCows: nCows, message: "there are no negative cows"}
	} else if nCows == 0 {
		return &InvalidCowsError{nCows: nCows, message: "no cows don't need food"}
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
