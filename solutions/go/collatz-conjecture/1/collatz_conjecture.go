// Package collatzconjecture
package collatzconjecture

import "errors"

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
