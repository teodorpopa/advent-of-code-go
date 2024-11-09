package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"math"
	"strings"
)

//go:embed input.txt
var input string

type Item struct {
	Name   string
	Cost   int
	Damage int
	Armor  int
}

var (
	weapons = []Item{}
	armors  = []Item{}
	rings   = []Item{}
)

func initShop() {
	weapons = append(weapons,
		Item{
			"Dagger",
			8,
			4,
			0,
		},
		Item{
			"Shortsword",
			10,
			5,
			0,
		},
		Item{
			"Warhamer",
			25,
			6,
			0,
		},
		Item{
			"Longsword",
			40,
			7,
			0,
		},
		Item{
			"Greataxe",
			74,
			8,
			0,
		},
	)

	armors = append(armors,
		Item{
			"Empty",
			0,
			0,
			0,
		},
		Item{
			"Leather",
			13,
			0,
			1,
		},
		Item{
			"Chainmail",
			31,
			0,
			2,
		},
		Item{
			"Splintmail",
			53,
			0,
			3,
		},
		Item{
			"Bandedmail",
			75,
			0,
			4,
		},
		Item{
			"Platemail",
			102,
			0,
			5,
		},
	)

	rings = append(rings,
		Item{
			"Empty",
			0,
			0,
			0,
		},
		Item{
			"Damage +1",
			25,
			1,
			0,
		},
		Item{
			"Damage +2",
			50,
			2,
			0,
		},
		Item{
			"Damage +3",
			100,
			3,
			0,
		},
		Item{
			"Defense +1",
			20,
			0,
			1,
		},
		Item{
			"Defense +2",
			40,
			0,
			2,
		},
		Item{
			"Defense +3",
			80,
			0,
			3,
		},
	)
}

func part1(input string) int {

	outfitCombinations := [][]Item{}
	bossHp, bossDamage, bossArmor := parseInput(input)
	playerHp := 100

	for _, weapon := range weapons {
		for _, armor := range armors {
			for _, ringLeft := range rings {
				for _, ringRight := range rings {
					c := []Item{}

					c = append(c, weapon, armor, ringLeft, ringRight)
					outfitCombinations = append(outfitCombinations, c)
				}
			}
		}
	}

	minCost := math.MaxInt32
	for _, outfitCombination := range outfitCombinations {
		playerWon, playerCost := simulateFight(outfitCombination, playerHp, bossHp, bossArmor, bossDamage)

		if playerWon == true {
			minCost = utils.Min(playerCost, minCost)
		}
	}

	return minCost
}

func part2(input string) int {
	outfitCombinations := [][]Item{}
	bossHp, bossDamage, bossArmor := parseInput(input)
	playerHp := 100

	for _, weapon := range weapons {
		for _, armor := range armors {
			for _, ringLeft := range rings {
				for _, ringRight := range rings {
					c := []Item{}

					c = append(c, weapon, armor, ringLeft, ringRight)
					outfitCombinations = append(outfitCombinations, c)
				}
			}
		}
	}

	maxCost := 0
	for _, outfitCombination := range outfitCombinations {
		playerWon, playerCost := simulateFight(outfitCombination, playerHp, bossHp, bossArmor, bossDamage)

		if playerWon == false {
			maxCost = utils.Max(playerCost, maxCost)
		}
	}

	return maxCost
}

func simulateFight(outfit []Item, playerHp int, bossHp int, bossArmor int, bossDamage int) (bool, int) {
	playerArmor := 0
	playerDamage := 0
	playerCost := 0

	for _, o := range outfit {
		playerArmor += o.Armor
		playerDamage += o.Damage
		playerCost += o.Cost
	}

	bossAttack := utils.Max(bossDamage-playerArmor, 1)
	playerAttack := utils.Max(playerDamage-bossArmor, 1)

	for bossHp >= 0 && playerHp >= 0 {
		bossHp -= playerAttack
		playerHp -= bossAttack
	}

	won := false
	if bossHp <= 0 {
		won = true
	}

	return won, playerCost
}

func parseInput(input string) (int, int, int) {
	lines := utils.ReadLines(input)

	hp := parseLine(lines[0])
	damage := parseLine(lines[1])
	armor := parseLine(lines[2])

	return hp, damage, armor
}

func parseLine(line string) int {
	parts := strings.Split(line, ": ")
	return utils.ToInt(parts[1])
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "which part tu run? 1 or 2")
	flag.Parse()

	initShop()

	res := 0
	if part == 1 {
		res = part1(input)
	} else {
		res = part2(input)
	}

	fmt.Println("Result: ", res)
}
