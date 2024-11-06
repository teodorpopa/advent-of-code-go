package main

import (
	"container/heap"
	_ "embed"
	"flag"
	"fmt"
	"image"
	"math"
	"strings"
)

//go:embed input.txt
var input string

type pqi[T any] struct {
	v T
	p int
}

type PQ[T any] []pqi[T]

func (q PQ[_]) Len() int {
	return len(q)
}

func (q PQ[_]) Less(i, j int) bool {
	return q[i].p < q[j].p
}

func (q PQ[_]) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *PQ[T]) Push(x any) {
	*q = append(*q, x.(pqi[T]))
}

func (q *PQ[_]) Pop() (x any) {
	x, *q = (*q)[len(*q)-1], (*q)[:len(*q)-1]
	return x
}

func (q *PQ[T]) GPush(v T, p int) {
	heap.Push(q, pqi[T]{v, p})
}

func (q *PQ[T]) GPop() (T, int) {
	x := heap.Pop(q).(pqi[T])
	return x.v, x.p
}

type State struct {
	Pos image.Point
	Dir image.Point
}

func run(input string, min, max int) int {
	grid, end := initGrid(input)

	queue := PQ[State]{}
	seen := map[State]struct{}{}

	queue.GPush(State{
		image.Point{
			0,
			0,
		},
		image.Point{
			1,
			0,
		},
	}, 0)

	queue.GPush(State{
		image.Point{
			0,
			0,
		},
		image.Point{
			0,
			1,
		},
	}, 0)

	for len(queue) > 0 {
		state, heat := queue.GPop()

		if state.Pos == end {
			return heat
		}
		if _, ok := seen[state]; ok {
			continue
		}
		seen[state] = struct{}{}

		for i := -max; i <= max; i++ {
			n := state.Pos.Add(state.Dir.Mul(i))
			if _, ok := grid[n]; !ok || i > -min && i < min {
				continue
			}
			h, s := 0, int(math.Copysign(1, float64(i)))
			for j := s; j != i+s; j += s {
				h += grid[state.Pos.Add(state.Dir.Mul(j))]
			}
			queue.GPush(State{
				n,
				image.Point{
					state.Dir.Y,
					state.Dir.X,
				},
			}, heat+h)
		}
	}
	return -1
}

func initGrid(input string) (map[image.Point]int, image.Point) {
	grid, end := map[image.Point]int{}, image.Point{0, 0}
	for y, s := range strings.Fields(input) {
		for x, r := range s {
			grid[image.Point{x, y}] = int(r - '0')
			end = image.Point{x, y}
		}
	}
	return grid, end
}

func part1(input string) int {
	return run(input, 1, 3)
}

func part2(input string) int {
	return run(input, 4, 10)
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
