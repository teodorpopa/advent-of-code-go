package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Number interface{ int | float64 }

func ToInt(i string) int {
	v, _ := strconv.Atoi(i)
	return v
}

func ToFloat(i string) float64 {
	return float64(ToInt(i))
}

func IsInt(i float64) bool {
	return i == float64(int(i))
}

func ToIntSlice(s []string) []int {
	var res []int
	for _, v := range s {
		res = append(res, ToInt(strings.TrimSpace(v)))
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

func AbsT[T Number](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

func ConcatTwoInt(i1, i2 int) int {
	num := fmt.Sprintf("%d%d", i1, i2)
	return ToInt(num)
}
