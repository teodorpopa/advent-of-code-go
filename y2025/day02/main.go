package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/teodorpopa/advent-of-code-go/utils"
)

//go:embed input.txt
var input string

func checkInvalidPid(pid string) int {
	if len(pid)%2 == 0 && pid[0:len(pid)/2] == pid[len(pid)/2:] {
		return utils.ToInt(pid)
	}
	return 0
}

func hasRepeatingPattern(pid string) int {
	n := len(pid)

	for i := 1; i <= n/2; i++ {
		if n%i != 0 {
			continue
		}

		pattern := pid[:i]
		valid := true

		for j := i; j < n; j += i {
			if pid[j:j+i] != pattern {
				valid = false
				break
			}
		}

		if valid {
			return utils.ToInt(pid)
		}
	}
	return 0
}

func part1(input string) int {
	parts := strings.Split(input, ",")
	sum := 0

	for _, part := range parts {
		pids := strings.Split(part, "-")

		from := utils.ToInt(pids[0])
		to := utils.ToInt(pids[1])

		for i := from; i <= to; i++ {
			sum += checkInvalidPid(utils.ToString(i))
		}
	}

	return sum
}

func part2(input string) int {
	parts := strings.Split(input, ",")
	sum := 0

	for _, part := range parts {
		pids := strings.Split(part, "-")

		from := utils.ToInt(pids[0])
		to := utils.ToInt(pids[1])

		for i := from; i <= to; i++ {
			sum += hasRepeatingPattern(utils.ToString(i))
		}
	}

	return sum
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
