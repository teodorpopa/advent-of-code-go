package main

import (
	_ "embed"
	"encoding/json"
	"flag"
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
