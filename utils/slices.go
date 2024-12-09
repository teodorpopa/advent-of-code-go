package utils

func CopyIntSlice(orig []int) []int {
	newNums := make([]int, len(orig))
	copy(newNums, orig[:])
	return newNums
}
