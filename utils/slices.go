package utils

func CopyIntSlice(orig []int) []int {
	newNums := make([]int, len(orig))
	copy(newNums, orig[:])
	return newNums
}

func ValidIndex(x, y, maxX, maxY int) bool {
	if x < 0 || y < 0 || x >= maxX || y >= maxY {
		return false
	}

	return true

}
