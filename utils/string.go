package utils

import (
	"sort"
	"strconv"
)

func ToString(s int) string {
	v := strconv.Itoa(s)
	return v
}

func SortString(s string) string {
	chars := []rune(s)
	sort.Slice(chars, func(i, j int) bool { //sort the string using the function
		return chars[i] < chars[j]
	})

	return string(chars)
}

func SortStringReverse(s string) string {
	chars := []rune(s)
	sort.Slice(chars, func(i, j int) bool { //sort the string using the function
		return chars[i] > chars[j]
	})

	return string(chars)
}

func ReverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
