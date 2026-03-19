// Package sorting processes anything that comes into a room by categorizing it with a label.
package sorting

import (
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return "This is the number " + strconv.FormatFloat(f, 'f', 1, 64)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return "This is a box containing the number " + strconv.FormatFloat(float64(nb.Number()), 'f', 1, 64)
}

type FancyNumberBox interface {
	Value() string
}

type FancyNumber struct {
	n string
}

// Value returns the numeric value of a FancyNumber.
func (i FancyNumber) Value() string {
	return i.n
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	value, ok := fnb.(FancyNumber)
	if !ok {
		return 0
	}
	number, err := strconv.Atoi(value.Value())
	if err != nil {
		return 0
	}
	return number
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	number := ExtractFancyNumber(fnb)
	if number == 0 {
		return "This is a fancy box containing the number 0.0"
	}
	return "This is a fancy box containing the number " + strconv.FormatFloat(float64(number), 'f', 1, 64)
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i any) string {
	result := "Return to sender"
	switch value := i.(type) {
	case int:
		result = DescribeNumber(float64(value))
	case float64:
		result = DescribeNumber(value)
	case NumberBox:
		result = DescribeNumberBox(value)
	case FancyNumberBox:
		result = DescribeFancyNumberBox(value)
	}
	return result
}
