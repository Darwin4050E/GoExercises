// Package collatzconjecture implements a function to return the number of steps CollatzConjeture needs to reach 1.
package collatzconjecture

import "errors"

// CollatzConjecture returns the number of steps it takes to reach 1 according to the rules of the Collatz Conjecture.
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("Negative or zero.")
	}
	count := 0
	result := n
	for result > 1 {
		if result%2 == 0 {
			result /= 2
		} else {
			result = 3*result + 1
		}
		count++
	}
	return count, nil
}
