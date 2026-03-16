// Package birdwatcher helps to keep track of how many birds have visited your garden.
package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	result := 0
	for index := 0; index < len(birdsPerDay); index++ {
		result += birdsPerDay[index]
	}
	return result
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	result := 0
	for index := (week - 1) * 7; index < week*7; index++ {
		result += birdsPerDay[index]
	}
	return result
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	result := birdsPerDay[:]
	for index := 0; index < len(result); index += 2 {
		result[index] += 1
	}
	return result
}
