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

type num uint16

type part struct {
	lhs []string
	rhs string
}

type wireVal struct {
	name  string
	value num
}

var binaryPart = map[string]func(inA, inB num) num{
	"AND":    func(inA, inB num) num { return num(inA & inB) },
	"OR":     func(inA, inB num) num { return num(inA | inB) },
	"LSHIFT": func(inA, inB num) num { return num(inA << inB) },
	"RSHIFT": func(inA, inB num) num { return num(inA >> inB) },
}

func getParts(input string) []part {
	lines := utils.ReadLines(input)
	parts := []part{}
	for _, line := range lines {
		tokens := strings.Split(line, " -> ")
		p := part{lhs: strings.Split(tokens[0], " "), rhs: tokens[1]}
		parts = append(parts, p)
	}
	return parts
}

func run(input string, overwrite map[string]int) int {
	parts := getParts(input)
	wire := map[string]chan num{}
	for _, p := range parts {
		wire[p.rhs] = make(chan num, 100)
	}

	done := make(chan bool)

	wireValues := map[string]num{}
	wireChan := make(chan wireVal)
	go func() {
		for i := 0; i < len(wire); i++ {
			wv := <-wireChan
			wireValues[wv.name] = wv.value
		}
		done <- true
	}()

	for _, rawP := range parts {
		p := rawP
		lhs, out := p.lhs, p.rhs
		if len(lhs) == 1 {
			n, err := strconv.Atoi(lhs[0])
			if err == nil {
				go func() {
					result := num(n)

					if overwrite[out] != 0 {
						result = num(overwrite[out])
					}

					wireChan <- wireVal{out, result}
					wire[out] <- result
				}()
			} else {
				go func() {
					result := <-wire[lhs[0]]
					wire[lhs[0]] <- result
					wireChan <- wireVal{out, result}
					wire[out] <- result
				}()
			}
		} else if len(lhs) == 2 {
			go func() {
				in := <-wire[lhs[1]]
				wire[lhs[1]] <- in
				result := ^in
				wireChan <- wireVal{out, result}
				wire[out] <- result
			}()
		} else {
			l, op, r := lhs[0], lhs[1], lhs[2]
			fn := binaryPart[op]
			go func() {
				inA, lok := wire[l]
				inB, rok := wire[r]
				var numA, numB num
				if lok && rok {
					numA = <-inA
					inA <- numA
					numB = <-inB
					inB <- numB
				} else if lok && !rok {
					b := utils.ToInt(r)
					numA = <-inA
					inA <- numA
					numB = num(b)
				} else if !lok && rok {
					a := utils.ToInt(l)
					numA = num(a)
					numB = <-inB
					inB <- numB
				} else {
					a := utils.ToInt(l)
					b := utils.ToInt(r)
					numA = num(a)
					numB = num(b)
				}
				result := fn(numA, numB)
				wireChan <- wireVal{out, result}
				wire[out] <- result
			}()
		}
	}

	<-done

	return int(wireValues["a"])
}

func part1(input string) int {
	overwrite := make(map[string]int)
	return run(input, overwrite)
}

func part2(input string) int {
	overwrite := map[string]int{
		"b": part1(input),
	}
	return run(input, overwrite)
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
