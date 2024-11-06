package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

const Approve = "A"
const Reject = "R"

const Lower = "<"
const Greater = ">"

const Min = 1
const Max = 4_000

type part struct {
	x int
	m int
	a int
	s int
}

func (p *part) getSum() int {
	return p.x + p.m + p.a + p.s
}

func parseParts(parts []string) []part {
	var pts = []part{}

	for _, p := range parts {
		p1 := strings.Trim(p, "{}")
		p2 := strings.Split(p1, "=")

		xp := strings.Split(p2[1], ",")
		mp := strings.Split(p2[2], ",")
		ap := strings.Split(p2[3], ",")
		sp := strings.Split(p2[4], ",")

		prt := part{
			x: utils.ToInt(xp[0]),
			m: utils.ToInt(mp[0]),
			a: utils.ToInt(ap[0]),
			s: utils.ToInt(sp[0]),
		}
		pts = append(pts, prt)
	}
	return pts
}

func parseWorkflows(wf []string) map[string][]string {
	var wfs = make(map[string][]string, len(wf))

	for _, w := range wf {
		pts := strings.Split(w, "{")
		cmp := strings.Trim(pts[1], "}")
		wfs[pts[0]] = strings.Split(cmp, ",")
	}

	return wfs
}

func check(p part, wf []string) string {
	var nextKey = ""
	var partKey = ""
	var value = 0
	var condition = ""

	for _, w := range wf {
		c := strings.Split(w, ":")

		if len(c) == 1 {
			return c[0]
		}
		nextKey = c[1]

		if strings.Contains(c[0], Lower) {
			lt := strings.Split(c[0], Lower)

			partKey = lt[0]
			value = utils.ToInt(lt[1])
			condition = Lower
		} else {
			gt := strings.Split(c[0], Greater)

			partKey = gt[0]
			value = utils.ToInt(gt[1])
			condition = Greater
		}

		cValue := 0
		if partKey == "x" {
			cValue = p.x
		} else if partKey == "m" {
			cValue = p.m
		} else if partKey == "a" {
			cValue = p.a
		} else if partKey == "s" {
			cValue = p.s
		}

		if condition == Lower {
			if cValue < value {
				return nextKey
			}
		} else if condition == Greater {
			if cValue > value {
				return nextKey
			}
		}

	}

	return ""
}

func parseWorkFlows(lines []string) map[string]string {
	var key string
	wfs := make(map[string]string)
	for _, line := range lines {
		if len(line) > 2 {
			ws := strings.Split(line[:len(line)-1], "{")
			if len(ws) != 2 {
				return wfs
			}
			ops := strings.Split(ws[1], ",")
			if len(ops) == 2 {
				key := ws[0]
				wfs[key] = ws[1]
			} else {
				for i := 0; i < len(ops)-1; i++ {
					key = generateKey(ws[0], i)
					wfs[key] = generateValue(ops, ops[i], ops[i+1], ws[0], i)
				}
			}
		} else {
			break
		}
	}
	return wfs
}

func generateKey(baseKey string, index int) string {
	if index == 0 {
		return baseKey
	}
	return fmt.Sprintf("%s%d", baseKey, index)
}

func generateValue(ops []string, op1, op2, baseKey string, index int) string {
	if index == len(ops)-2 {
		return fmt.Sprintf("%s,%s", op1, op2)
	}
	return fmt.Sprintf("%s,%s%d", op1, baseKey, index+1)
}

func parseCondition(condition string) (param byte, symbol string, value int) {
	parts := strings.Split(condition, ":")
	if len(parts) != 2 {
		return
	}
	param = parts[0][0]
	symbol = string(parts[0][1])
	value, _ = strconv.Atoi(parts[0][2:])
	return
}

func getOptions(condition string) (trueOption, falseOption string) {
	options := strings.Split(strings.Split(condition, ":")[1], ",")
	return options[0], options[1]
}

func calculatePossibleValues(wfs map[string]string, wfKey string, xMin, xMax, mMin, mMax, aMin, aMax, sMin, sMax int) int {
	if wfKey == Approve {
		return (xMax - xMin + 1) * (mMax - mMin + 1) * (aMax - aMin + 1) * (sMax - sMin + 1)
	} else if wfKey == Reject {
		return 0
	} else {
		acc := 0
		condition := wfs[wfKey]
		param, symbol, value := parseCondition(condition)
		trueOption, falseOption := getOptions(condition)
		switch param {
		case 'x':
			if symbol == Greater {
				acc += calculatePossibleValues(wfs, trueOption, value+1, xMax, mMin, mMax, aMin, aMax, sMin, sMax)
				acc += calculatePossibleValues(wfs, falseOption, xMin, value, mMin, mMax, aMin, aMax, sMin, sMax)
			} else {
				acc += calculatePossibleValues(wfs, trueOption, xMin, value-1, mMin, mMax, aMin, aMax, sMin, sMax)
				acc += calculatePossibleValues(wfs, falseOption, value, xMax, mMin, mMax, aMin, aMax, sMin, sMax)
			}
		case 'm':
			if symbol == Greater {
				acc += calculatePossibleValues(wfs, trueOption, xMin, xMax, value+1, mMax, aMin, aMax, sMin, sMax)
				acc += calculatePossibleValues(wfs, falseOption, xMin, xMax, mMin, value, aMin, aMax, sMin, sMax)
			} else {
				acc += calculatePossibleValues(wfs, trueOption, xMin, xMax, mMin, value-1, aMin, aMax, sMin, sMax)
				acc += calculatePossibleValues(wfs, falseOption, xMin, xMax, value, mMax, aMin, aMax, sMin, sMax)
			}
		case 'a':
			if symbol == Greater {
				acc += calculatePossibleValues(wfs, trueOption, xMin, xMax, mMin, mMax, value+1, aMax, sMin, sMax)
				acc += calculatePossibleValues(wfs, falseOption, xMin, xMax, mMin, mMax, aMin, value, sMin, sMax)
			} else {
				acc += calculatePossibleValues(wfs, trueOption, xMin, xMax, mMin, mMax, aMin, value-1, sMin, sMax)
				acc += calculatePossibleValues(wfs, falseOption, xMin, xMax, mMin, mMax, value, aMax, sMin, sMax)
			}
		case 's':
			if symbol == Greater {
				acc += calculatePossibleValues(wfs, trueOption, xMin, xMax, mMin, mMax, aMin, aMax, value+1, sMax)
				acc += calculatePossibleValues(wfs, falseOption, xMin, xMax, mMin, mMax, aMin, aMax, sMin, value)
			} else {
				acc += calculatePossibleValues(wfs, trueOption, xMin, xMax, mMin, mMax, aMin, aMax, sMin, value-1)
				acc += calculatePossibleValues(wfs, falseOption, xMin, xMax, mMin, mMax, aMin, aMax, value, sMax)
			}
		}
		return acc
	}
}

func part1(input string) int {
	blocks := strings.Split(input, "\n\n")

	wf := strings.Split(blocks[0], "\n")
	wfs := parseWorkflows(wf)

	parts := strings.Split(blocks[1], "\n")
	pts := parseParts(parts)

	sum := 0

	for _, p := range pts {
		startKey := "in"
		actionFound := false

		for {
			result := check(p, wfs[startKey])

			if result == Approve {
				sum += p.getSum()
				actionFound = true
			} else if result == Reject {
				actionFound = true
			} else {
				startKey = result
			}

			if actionFound == true {
				break
			}
		}
	}

	return sum
}

func part2(input string) int {
	blocks := strings.Split(input, "\n\n")

	wf := strings.Split(blocks[0], "\n")
	wfs := parseWorkFlows(wf)

	sum := calculatePossibleValues(
		wfs,
		"in",
		Min,
		Max,
		Min,
		Max,
		Min,
		Max,
		Min,
		Max,
	)

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
