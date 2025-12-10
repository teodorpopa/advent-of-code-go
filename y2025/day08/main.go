package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/teodorpopa/advent-of-code-go/utils"
)

//go:embed input.txt
var input string

type coordinates struct {
	x, y, z int
}

type box struct {
	coordinates coordinates
	circuit     *circuit
}

type circuit struct {
	boxes []*box
}

type edge struct {
	box1, box2 *box
	distance   int
}

func getBoxes(input string) []*box {
	boxes := []*box{}
	for _, line := range utils.ReadLines(input) {
		parts := strings.Split(line, ",")
		boxes = append(boxes, &box{
			coordinates: coordinates{
				x: utils.ToInt(parts[0]),
				y: utils.ToInt(parts[1]),
				z: utils.ToInt(parts[2]),
			},
		})
	}
	return boxes
}

func calcDistance(b1, b2 *box) int {
	x := b1.coordinates.x - b2.coordinates.x
	y := b1.coordinates.y - b2.coordinates.y
	z := b1.coordinates.z - b2.coordinates.z
	return x*x + y*y + z*z
}

func getEdges(boxes []*box, distances map[*box]map[*box]int) []edge {
	edges := []edge{}
	for i, b1 := range boxes {
		distances[b1] = map[*box]int{}
		for _, b2 := range boxes[i+1:] {
			edges = append(edges, edge{
				box1:     b1,
				box2:     b2,
				distance: calcDistance(b1, b2),
			})
		}
	}

	return edges
}

func part1(input string, iterations int) int {

	boxes := getBoxes(input)
	distances := map[*box]map[*box]int{}
	allEdges := getEdges(boxes, distances)
	circuits := []*circuit{}
	limit := 3

	sort.Slice(allEdges, func(i int, j int) bool {
		return allEdges[i].distance < allEdges[j].distance
	})

	for range iterations {
		thisEdge := allEdges[0]
		allEdges = allEdges[1:]

		if thisEdge.box1.circuit == nil && thisEdge.box2.circuit == nil {
			circuits = append(circuits, &circuit{
				boxes: []*box{
					thisEdge.box1,
					thisEdge.box2,
				},
			})
			thisEdge.box1.circuit = circuits[len(circuits)-1]
			thisEdge.box2.circuit = circuits[len(circuits)-1]
		} else if thisEdge.box1.circuit == nil {
			thisEdge.box2.circuit.boxes = append(thisEdge.box2.circuit.boxes, thisEdge.box1)
			thisEdge.box1.circuit = thisEdge.box2.circuit
		} else if thisEdge.box2.circuit == nil {
			thisEdge.box1.circuit.boxes = append(thisEdge.box1.circuit.boxes, thisEdge.box2)
			thisEdge.box2.circuit = thisEdge.box1.circuit
		} else if thisEdge.box1.circuit != thisEdge.box2.circuit {
			oldCircuit := thisEdge.box2.circuit

			for _, box := range thisEdge.box2.circuit.boxes {
				thisEdge.box1.circuit.boxes = append(thisEdge.box1.circuit.boxes, box)
				box.circuit = thisEdge.box1.circuit
			}

			for i := range circuits {
				if circuits[i] == oldCircuit {
					circuits[len(circuits)-1], circuits[i] = circuits[i], circuits[len(circuits)-1]
					circuits = circuits[:len(circuits)-1]
					break
				}
			}
		}
	}

	circuitSizes := []int{}
	for _, c := range circuits {
		circuitSizes = append(circuitSizes, len(c.boxes))
	}

	sort.Ints(circuitSizes)

	mult := 1
	for i := range limit {
		mult *= circuitSizes[len(circuitSizes)-1-i]
	}
	return mult
}

func part2(input string) int {

	boxes := getBoxes(input)
	distances := map[*box]map[*box]int{}
	allEdges := getEdges(boxes, distances)
	circuits := []*circuit{}

	sort.Slice(allEdges, func(i int, j int) bool {
		return allEdges[i].distance < allEdges[j].distance
	})

	for len(allEdges) > 0 {
		thisEdge := allEdges[0]
		allEdges = allEdges[1:]

		if thisEdge.box1.circuit == nil && thisEdge.box2.circuit == nil {
			circuits = append(circuits, &circuit{
				boxes: []*box{
					thisEdge.box1,
					thisEdge.box2,
				},
			})
			thisEdge.box1.circuit = circuits[len(circuits)-1]
			thisEdge.box2.circuit = circuits[len(circuits)-1]
		} else if thisEdge.box1.circuit == nil {
			thisEdge.box2.circuit.boxes = append(thisEdge.box2.circuit.boxes, thisEdge.box1)
			thisEdge.box1.circuit = thisEdge.box2.circuit
		} else if thisEdge.box2.circuit == nil {
			thisEdge.box1.circuit.boxes = append(thisEdge.box1.circuit.boxes, thisEdge.box2)
			thisEdge.box2.circuit = thisEdge.box1.circuit
		} else if thisEdge.box1.circuit != thisEdge.box2.circuit {
			oldCircuit := thisEdge.box2.circuit

			for _, box := range thisEdge.box2.circuit.boxes {
				thisEdge.box1.circuit.boxes = append(thisEdge.box1.circuit.boxes, box)
				box.circuit = thisEdge.box1.circuit
			}

			for i := range circuits {
				if circuits[i] == oldCircuit {
					circuits[len(circuits)-1], circuits[i] = circuits[i], circuits[len(circuits)-1]
					circuits = circuits[:len(circuits)-1]
					break
				}
			}
		}

		if len(thisEdge.box1.circuit.boxes) == len(boxes) {
			return thisEdge.box1.coordinates.x * thisEdge.box2.coordinates.x
		}
	}

	return 0
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "which part tu run? 1 or 2")
	flag.Parse()

	res := 0
	if part == 1 {
		res = part1(input, 1000)
	} else {
		res = part2(input)
	}

	fmt.Println("Result: ", res)
}
