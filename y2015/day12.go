package y2015

import (
	"encoding/json"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"reflect"
	"strconv"
)

func readJson(input string) []interface{} {
	var m []interface{}

	err := json.Unmarshal([]byte(input), &m)
	if err != nil {
		panic(err)
	}

	return m
}

func sumJson(json []interface{}, ignoreRed bool) int {
	var sum float64

	for _, v := range json {
		sum += calcSum(v, ignoreRed)
	}

	return int(sum)
}

func calcSum(v interface{}, ignoreRed bool) float64 {
	var sum float64

	if reflect.TypeOf(v).Kind() == reflect.Slice || reflect.TypeOf(v).Kind() == reflect.Array {
		for _, val := range v.([]interface{}) {
			sum += calcSum(val, ignoreRed)
		}
	} else if reflect.TypeOf(v).Kind() == reflect.Map {
		if ignoreRed {
			for _, val := range v.(map[string]interface{}) {
				if reflect.TypeOf(val).Kind() == reflect.String && val.(string) == "red" {
					return 0
				}
			}
		}

		for _, val := range v.(map[string]interface{}) {
			sum += calcSum(val, ignoreRed)
		}
	} else if reflect.TypeOf(v).Kind() == reflect.Float64 {
		sum += v.(float64)
	} else if nr, err := strconv.Atoi(v.(string)); err == nil {
		sum += float64(nr)
	}

	return sum
}

func Day12First(input string) int {
	return sumJson(readJson(input), false)
}

func Day12Second(input string) int {
	return sumJson(readJson(input), true)
}

func Day12() {
	fmt.Println(utils.DAY_PREFIX, "Day 12")
	fmt.Println(utils.PART_PREFIX, "Part 1:", Day12First(utils.ReadFile("y2015/input/day12.txt")))
	fmt.Println(utils.PART_PREFIX, "Part 2:", Day12Second(utils.ReadFile("y2015/input/day12.txt")))
}
