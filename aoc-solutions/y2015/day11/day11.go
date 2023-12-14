package day11

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

var start, end = byte(97), byte(122)
var ignore = []rune{'i', 'o', 'l'}

func noIgnores(password string) bool {
	ret := true
	for i := 0; i < len(ignore); i++ {
		ret = ret && strings.IndexByte(password, byte(ignore[i])) == -1
	}
	return ret
}

func hasConsecutive(password string) bool {
	for i := 0; i < len(password)-2; i++ {
		if password[i]+1 == password[i+1] && password[i]+2 == password[i+2] {
			return true
		}
	}
	return false
}

func twoPairs(password string) bool {
	pairs := 0
	for i := 0; i < len(password)-1; i++ {
		if password[i] == password[i+1] {
			pairs++
			i++
		}
	}
	return pairs >= 2
}

func increment(password string) string {
	return incRec(password, len(password)-1)
}

func incRec(password string, ch int) string {
	cp := password
	b := cp[ch]
	b = b + 1

	if loop := b > end; loop {
		b = start
		cp = incRec(cp, ch-1)
	}
	cp = cp[0:ch] + string(b) + cp[ch+1:]
	return cp
}

func part1(input string) string {
	password := input

	for {
		password = increment(password)
		if hasConsecutive(password) && noIgnores(password) && twoPairs(password) {
			return password
		}
	}
}

func part2(input string) string {
	newInput := part1(input)
	return part1(newInput)
}

func Solve(part int) int {
	if part == 1 {
		fmt.Println(part1(input))
		return 0
	} else if part == 2 {
		fmt.Println(part2(input))
		return 0
	} else {
		fmt.Println("Invalid part to run")
		return -1
	}
}
