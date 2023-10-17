package y2015

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"strconv"
	"strings"
)

func md5string(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func checkString(hash string, prefix string) bool {
	return strings.HasPrefix(hash, prefix)
}

func Day04First(input string) int {

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

func Day04Second(input string) int {

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

func Day04() {
	fmt.Println(utils.DAY_PREFIX, "Day 04")
	fmt.Println(utils.PART_PREFIX, "Part 1:", Day04First(utils.ReadFile("y2015/input/day04.txt")))
	fmt.Println(utils.PART_PREFIX, "Part 2:", Day04Second(utils.ReadFile("y2015/input/day04.txt")))
}
