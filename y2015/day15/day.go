package day15

import (
	_ "embed"
	"fmt"
	"regexp"

	"github.com/teodorpopa/advent-of-code-go/utils"
)

//go:embed input.txt
var input string

type Ingredient struct {
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func part1(input string) int {
	lines := utils.ReadLines(input)
	ingredients := getIngredients(lines)
	score := calculateProperties(ingredients, false)

	return score
}

func part2(input string) int {
	lines := utils.ReadLines(input)
	ingredients := getIngredients(lines)
	score := calculateProperties(ingredients, true)

	return score
}

func getIngredients(lines []string) []Ingredient {
	r, _ := regexp.Compile("(.*): capacity (.*), durability (.*), flavor (.*), texture (.*), calories (.*)")
	ingredients := []Ingredient{}

	for _, l := range lines {
		founds := r.FindStringSubmatch(l)

		ingredient := Ingredient{
			capacity:   utils.ToInt(founds[2]),
			durability: utils.ToInt(founds[3]),
			flavor:     utils.ToInt(founds[4]),
			texture:    utils.ToInt(founds[5]),
			calories:   utils.ToInt(founds[6]),
		}

		ingredients = append(ingredients, ingredient)

	}
	return ingredients
}

func calculateProperties(ingredients []Ingredient, cals bool) int {
	score := 0

	for ing1 := 0; ing1 < 100; ing1++ {
		for ing2 := 0; ing2 < 100; ing2++ {
			for ing3 := 0; ing3 < 100; ing3++ {
				ing4 := 100 - ing1 - ing2 - ing3

				if ing4 > 0 {
					cap := ing1*ingredients[0].capacity + ing2*ingredients[1].capacity + ing3*ingredients[2].capacity + ing4*ingredients[3].capacity
					dur := ing1*ingredients[0].durability + ing2*ingredients[1].durability + ing3*ingredients[2].durability + ing4*ingredients[3].durability
					fla := ing1*ingredients[0].flavor + ing2*ingredients[1].flavor + ing3*ingredients[2].flavor + ing4*ingredients[3].flavor
					tex := ing1*ingredients[0].texture + ing2*ingredients[1].texture + ing3*ingredients[2].texture + ing4*ingredients[3].texture

					calories := ing1*ingredients[0].calories + ing2*ingredients[1].calories + ing3*ingredients[2].calories + ing4*ingredients[3].calories

					cap = utils.Max(0, cap)
					dur = utils.Max(0, dur)
					fla = utils.Max(0, fla)
					tex = utils.Max(0, tex)

					newScore := cap * dur * fla * tex

					if cals {
						if calories == 500 {
							score = utils.Max(newScore, score)
						}
					} else {
						score = utils.Max(newScore, score)
					}
				}
			}
		}
	}

	return score
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
