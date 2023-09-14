package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

var files = []string{"A", "B", "C", "D", "E", "F", "G", "H", "Z"}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var count int
	for _, square := range cb[file] {
		if square {
			count += 1
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	var count int
	if rank > 8 || rank < 1 {
		return count
	}
	index := rank - 1

	for _, file := range cb {
		if file[index] {
			count += 1
		}

	}

	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var count int
	for _, file := range cb {
		count += len(file)
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var count int
	for file := range cb {
		count += CountInFile(cb, file)
	}
	return count
}
