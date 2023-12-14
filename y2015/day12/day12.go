package day12

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

//go:embed input.txt
var input string

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

func part1(input string) int {
	return sumJson(readJson(input), false)
}

func part2(input string) int {
	return sumJson(readJson(input), true)
}

func Solve(part int) int {
	if part == 1 {
		return part1(input)
	} else if part == 2 {
		return part2(input)
	} else {
		fmt.Println("Invalid part to run")
		return -1
	}
}
