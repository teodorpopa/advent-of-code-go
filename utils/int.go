package utils

import (
	"math"
	"strconv"
)

func ToInt(i string) int {
	v, _ := strconv.Atoi(i)
	return v
}

func ToIntSlice(s []string) []int {
	var res []int
	for _, v := range s {
		res = append(res, ToInt(v))
	}
	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(x int) int {
	return int(math.Abs(float64(x)))
}
