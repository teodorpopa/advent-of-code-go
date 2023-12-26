package day25

import (
	_ "embed"
	"fmt"
	"github.com/emirpasic/gods/queues/priorityqueue"
	"strings"
)

//go:embed input.txt
var input string

type component struct {
	name  string
	edges map[string]*component
}

type node struct {
	component *component
	dist      int
}

func getComponents(input string) map[string]*component {
	lines := strings.Split(input, "\n")
	components := make(map[string]*component)
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		name := parts[0]
		c, ok := components[name]
		if !ok {
			c = &component{
				name:  name,
				edges: make(map[string]*component),
			}
			components[name] = c
		}

		connections := strings.Split(parts[1], " ")
		for _, connection := range connections {
			conn, ok := components[connection]
			if !ok {
				conn = &component{
					name:  connection,
					edges: make(map[string]*component),
				}
				components[connection] = conn
			}
			c.edges[connection] = conn
			conn.edges[name] = c
		}
	}
	return components
}

func getDistance(from, to *component, components map[string]*component) int {
	distances := make(map[*component]int)
	visited := make(map[*component]bool)
	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		return a.(*node).dist - b.(*node).dist
	})

	for _, c := range components {
		distances[c] = int(^uint(0) >> 1)
	}
	distances[from] = 0

	pq.Enqueue(&node{
		component: from,
		dist:      0,
	})

	for !pq.Empty() {
		v, _ := pq.Dequeue()
		n := v.(*node)
		current := n.component

		if visited[current] {
			continue
		}
		visited[current] = true

		for _, connection := range current.edges {
			if current == from && connection == to {
				continue
			}
			distance := distances[current] + 1
			if distance < distances[connection] {
				distances[connection] = distance
				pq.Enqueue(&node{
					component: connection,
					dist:      distance,
				})
			}
		}
	}

	return distances[to]
}

func getGroupLength(start string, components map[string]*component) int {
	visited := make(map[*component]bool)
	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		return a.(*node).dist - b.(*node).dist
	})

	pq.Enqueue(&node{
		component: components[start], dist: 0,
	})

	for !pq.Empty() {
		v, _ := pq.Dequeue()
		n := v.(*node)
		current := n.component

		if visited[current] {
			continue
		}
		visited[current] = true

		for _, connection := range current.edges {
			pq.Enqueue(&node{
				component: connection,
				dist:      n.dist + 1,
			})
		}
	}

	return len(visited)
}

func part1(input string) int {
	components := getComponents(input)
	distances := make(map[string]int)

	for _, component := range components {
		for _, connection := range component.edges {
			if _, ok := distances[connection.name+component.name]; ok {
				continue
			}
			distances[component.name+connection.name] = getDistance(component, connection, components)
		}
	}

	highestThree := make([]string, 3)
	highestThreeDistances := make([]int, 3)
	for name, distance := range distances {
		for i, highestDistance := range highestThreeDistances {
			if distance > highestDistance {
				highestThree[i] = name
				highestThreeDistances[i] = distance
				break
			}
		}
	}

	for _, name := range highestThree {
		name1 := name[0:3]
		name2 := name[3:6]
		delete(components[name1].edges, name2)
		delete(components[name2].edges, name1)
	}

	group1Start := highestThree[0][0:3]
	group2Start := highestThree[0][3:6]

	return getGroupLength(group1Start, components) * getGroupLength(group2Start, components)
}

func Solve(part int) int {
	if part == 1 {
		return part1(input)
	} else if part == 2 {
		return 0
	} else {
		fmt.Println("Invalid part to run")
		return -1
	}
}
