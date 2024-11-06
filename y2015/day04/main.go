package main

import (
	"crypto/md5"
	_ "embed"
	"encoding/hex"
	"flag"
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

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "which part tu run? 1 or 2")
	flag.Parse()

	res := 0
	if part == 1 {
		res = part1(input)
	} else {
		res = part2(input)
	}

	fmt.Println("Result: ", res)
}
