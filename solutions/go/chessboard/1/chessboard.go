// Package chessboard implements a version of the Chess game where:
// The horizontal rows of squares, called ranks, are numbered 1 through 8.
// The vertical columns of squares, called files, are labeled A through H.
package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	col, exists := cb[file]
	count := 0
	if !exists {
		return count
	}
	for _, value := range col {
		if value {
			count++
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	count := 0
	if rank < 1 || rank > 8 {
		return count
	}
	for _, value := range cb {
		if value[rank-1] {
			count++
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0
	for range cb {
		count++
	}
	return count * 8
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for key := range cb {
		count += CountInFile(cb, key)
	}
	return count
}
