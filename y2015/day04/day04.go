package day04

import (
	"crypto/md5"
	_ "embed"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func md5string(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func checkString(hash string, prefix string) bool {
	return strings.HasPrefix(hash, prefix)
}

func part1(input string) int {

	i := 0
	for {
		string := input + strconv.Itoa(i)
		hash := md5string(string)

		if checkString(hash, "00000") {
			break
		}

		i++
	}

	return i
}

func part2(input string) int {

	i := 0
	for {
		string := input + strconv.Itoa(i)
		hash := md5string(string)

		if checkString(hash, "000000") {
			break
		}

		i++
	}

	return i
}

func Solve(part int) int {
	if part == 1 {
		return part1(input)
	} else if part == 2 {
		return part2(input)
	} else {
		fmt.Println("Invalid part to run")
		return -1
	}
}
