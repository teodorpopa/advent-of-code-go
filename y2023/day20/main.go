package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"strings"
)

//go:embed input.txt
var input string

type ModuleType int
type Pulse bool

const (
	UNTYPED ModuleType = iota
	BROADCASTER
	FLIP_FLOP
	CONJUNCTION
)

const (
	HIGH Pulse = true
	LOW  Pulse = false
)

const (
	PUSHES       = 1000
	FINAL_MODULE = "rx"
)

type Module struct {
	Type     ModuleType
	Dest     []string
	On       bool
	Remember map[string]Pulse
}

type Input struct {
	Pulse       Pulse
	Source      string
	Destination string
}

type Metadata struct {
	Targets        map[string]bool
	Counters       map[string]int
	ExpectedOutput Pulse
	PressCount     int
}

func parseInputs(input string) map[string]Module {
	lines := utils.ReadLines(input)
	modules := make(map[string]Module)

	for _, line := range lines {
		data := strings.Split(line, " -> ")
		var label string
		var moduleType ModuleType

		switch data[0][0] {
		case '%':
			label = data[0][1:]
			moduleType = FLIP_FLOP
		case '&':
			label = data[0][1:]
			moduleType = CONJUNCTION
		default:
			label = data[0]
			moduleType = UNTYPED
			if label == "broadcaster" {
				moduleType = BROADCASTER
			}
		}

		destinations := strings.Split(data[1], ", ")
		module, ok := modules[label]
		if ok {
			module.Type = moduleType
			module.Dest = destinations
			modules[label] = module
		} else {
			modules[label] = Module{
				Type:     moduleType,
				Dest:     destinations,
				Remember: make(map[string]Pulse),
			}
		}

		for _, dest := range destinations {
			destModule, ok := modules[dest]
			if ok {
				destModule.Remember[label] = LOW
				modules[dest] = destModule
			} else {
				modules[dest] = Module{
					Remember: map[string]Pulse{label: LOW},
				}
			}
		}
	}

	return modules
}

func isDefaultState(modules map[string]Module) bool {
	for _, module := range modules {
		if module.On {
			return false
		}

		for _, memory := range module.Remember {
			if memory == HIGH {
				return false
			}
		}
	}

	return true
}

func sendPulse(pulseCounter map[Pulse]int64, modules map[string]Module, inputs []Input, metadata *Metadata) {
	newInputs := []Input{}
	for _, input := range inputs {
		pulseCounter[input.Pulse]++

		destModule := modules[input.Destination]
		switch destModule.Type {
		case BROADCASTER:
			for _, newDest := range destModule.Dest {
				newInputs = append(newInputs, Input{
					Pulse:       input.Pulse,
					Source:      input.Destination,
					Destination: newDest,
				})
			}
		case FLIP_FLOP:
			if input.Pulse == LOW {
				var nextPulse Pulse
				if destModule.On {
					nextPulse = LOW
					destModule.On = false
				} else {
					nextPulse = HIGH
					destModule.On = true
				}
				modules[input.Destination] = destModule

				for _, newDest := range destModule.Dest {
					newInputs = append(newInputs, Input{
						Pulse:       nextPulse,
						Source:      input.Destination,
						Destination: newDest,
					})
				}
			}
		case CONJUNCTION:
			destModule.Remember[input.Source] = input.Pulse
			modules[input.Destination] = destModule

			nextPulse := LOW
			for _, memory := range destModule.Remember {
				if memory == LOW {
					nextPulse = HIGH
					break
				}
			}

			for _, newDest := range destModule.Dest {
				if metadata != nil {
					if _, ok := metadata.Targets[input.Destination]; ok && nextPulse == metadata.ExpectedOutput {
						metadata.Counters[input.Destination] = metadata.PressCount
						delete(metadata.Targets, input.Destination)
					}

					if len(metadata.Targets) == 0 {
						return
					}
				}

				newInputs = append(newInputs, Input{
					Pulse:       nextPulse,
					Source:      input.Destination,
					Destination: newDest,
				})
			}
		}
	}

	if len(newInputs) > 0 {
		sendPulse(pulseCounter, modules, newInputs, metadata)
	}
}

func part1(input string) int {
	modules := parseInputs(input)
	pulses := make(map[Pulse]int64)
	maxCount := PUSHES

	for c := 1; c <= maxCount; c++ {
		sendPulse(
			pulses,
			modules,
			[]Input{
				{
					Pulse:       LOW,
					Source:      "button",
					Destination: "broadcaster",
				},
			},
			nil,
		)

		if isDefaultState(modules) {
			multiplier := int64(maxCount / c)
			for pulse, amount := range pulses {
				pulses[pulse] = amount * multiplier
			}

			c = maxCount - (maxCount % c)
		}
	}

	return int(pulses[HIGH] * pulses[LOW])
}

func part2(input string) int {
	modules := parseInputs(input)
	finalModule := modules[FINAL_MODULE]

	var peModule Module
	for moduleLabel := range finalModule.Remember {
		peModule = modules[moduleLabel]
		break
	}

	metadataTargets := map[string]bool{}
	for moduleLabel := range peModule.Remember {
		metadataTargets[moduleLabel] = true
	}

	metadata := Metadata{
		Targets:        metadataTargets,
		Counters:       make(map[string]int),
		ExpectedOutput: HIGH,
		PressCount:     1,
	}

	for {
		sendPulse(
			make(map[Pulse]int64),
			modules,
			[]Input{{
				Pulse:       LOW,
				Source:      "button",
				Destination: "broadcaster",
			}},
			&metadata,
		)

		if len(metadata.Targets) == 0 {
			break
		}

		metadata.PressCount++
	}

	numbers := []int{}
	for _, count := range metadata.Counters {
		numbers = append(numbers, count)
	}
	return utils.LCM(numbers)
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
