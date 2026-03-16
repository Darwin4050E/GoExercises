// Package lasagnamaster prepares and cooks your brilliant lasagna from your favorite cookbook.
package lasagnamaster

import "slices"

// TODO: define the 'PreparationTime()' function
// PreparationTime calculates the estimate for the total preparation time based on the number of layers.
func PreparationTime(layers []string, preparation_time int) int {
	n_layers := len(layers)
	time := 2
	if preparation_time != 0 {
		time = preparation_time
	}
	return n_layers * time
}

// TODO: define the 'Quantities()' function
// Quantities determines the quantity of noodles and sauce needed to make your meal.
func Quantities(layers []string) (int, float64) {
	q_noodles := 0
	q_sauces := 0
	for index := 0; index < len(layers); index++ {
		switch layers[index] {
		case "noodles":
			q_noodles += 1
		case "sauce":
			q_sauces += 1
		}
	}
	return q_noodles * 50, float64(q_sauces) * 0.2
}

// TODO: define the 'AddSecretIngredient()' function
// AddSecretIngredient replaces your ingredient list with the last item from your friends list.
func AddSecretIngredient(friend_list, my_list []string) {
	my_list[len(my_list)-1] = friend_list[len(friend_list)-1]
}

// TODO: define the 'ScaleRecipe()' function
// ScaleRecipe calculates the amounts needed for the desired number of portions.
// It keeps the original recipe thought.
func ScaleRecipe(amounts []float64, portions int) []float64 {
	result := slices.Clone(amounts)
	for index := 0; index < len(result); index++ {
		result[index] = (result[index] / 2) * float64(portions)
	}
	return result
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
