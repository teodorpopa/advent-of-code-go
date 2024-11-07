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
