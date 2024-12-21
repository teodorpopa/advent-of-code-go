package utils

import (
	"golang.org/x/exp/rand"
	"sort"
	"time"
)

func Permutations(arr []interface{}) [][]interface{} {
	var helper func([]interface{}, int)
	var res [][]interface{}

	helper = func(arr []interface{}, n int) {
		if n == 1 {
			tmp := make([]interface{}, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func LCM(numbers []int) int {
	lcm := numbers[0]
	for i := 0; i < len(numbers); i++ {
		num1 := lcm
		num2 := numbers[i]
		gcd := 1
		for num2 != 0 {
			temp := num2
			num2 = num1 % num2
			num1 = temp
		}
		gcd = num1
		lcm = (lcm * numbers[i]) / gcd
	}

	return lcm
}

var rn = rand.New(rand.NewSource(uint64(time.Now().UnixNano())))

func ShuffleSlice(in []string) []string {
	sort.Slice(in, func(i, j int) bool {
		return rn.Intn(2) == 1
	})
	return in
}

func DFS(grid [][]string, visited [][]bool, i, j int, sides *[][3]int) (int, int) {
	visited[i][j] = true
	neighbors := GetNeighbors(i, j)
	perimeter := 0
	area := 1
	for _, n := range neighbors {
		ni, nj := n[0], n[1]
		if ni < 0 || ni >= len(grid) || nj < 0 || nj >= len(grid[ni]) || grid[ni][nj] != grid[i][j] {
			if sides != nil {
				*sides = append(*sides, [3]int{ni, nj, n[2]})
			}
			perimeter++
		} else if !visited[ni][nj] {
			a, p := DFS(grid, visited, ni, nj, sides)
			area += a
			perimeter += p
		}
	}

	return area, perimeter
}

func GetNeighbors(i, j int) [][3]int {
	return [][3]int{
		{i - 1, j, 0},
		{i + 1, j, 1},
		{i, j - 1, 2},
		{i, j + 1, 3},
	}
}
