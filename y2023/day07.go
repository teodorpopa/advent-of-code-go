package y2023

import (
	"fmt"
	"github.com/teodorpopa/advent-of-code-go/utils"
	"golang.org/x/exp/slices"
	"strings"
)

type Hand struct {
	Cards string
	Score int
	Hand  int
}

const (
	Five     int = 7
	Four         = 6
	Full         = 5
	Three        = 4
	TwoPairs     = 3
	OnePair      = 2
	High         = 1
)

func getHandType(cards string) int {
	cGroups := map[rune]int{}

	for _, card := range cards {
		if _, ok := cGroups[card]; ok {
			cGroups[card] += 1
		} else {
			cGroups[card] = 1
		}
	}

	for _, group := range cGroups {
		if group == 5 {
			return Five
		}
		if group == 4 {
			return Four
		}
	}

	if len(cGroups) == 2 {
		return Full
	}

	for _, group := range cGroups {
		if group == 3 {
			return Three
		}
	}

	pairs := 0
	for _, group := range cGroups {
		if group == 2 {
			pairs++
		}
	}
	if pairs == 2 {
		return TwoPairs
	} else if pairs == 1 {
		return OnePair
	}

	return High
}

func getHandTypePart2(cards string) int {
	cGroups := map[rune]int{}

	for _, card := range cards {
		if _, ok := cGroups[card]; ok {
			cGroups[card] += 1
		} else {
			cGroups[card] = 1
		}
	}

	jokers := cGroups['J']
	if jokers == 0 {
		return getHandType(cards)
	}

	var mk rune
	var mv int

	for k, v := range cGroups {
		if k == 'J' {
			continue
		}
		if v > mv {
			mk = k
			mv = v
		}
	}

	cards = strings.ReplaceAll(cards, "J", string(mk))
	return getHandType(cards)
}

func compareCards(a, b rune, part int) int {
	cmp := []rune{}

	if part == 1 {
		cmp = []rune{
			'A',
			'K',
			'Q',
			'J',
			'T',
			'9',
			'8',
			'7',
			'6',
			'5',
			'4',
			'3',
			'2',
		}
	} else {
		cmp = []rune{
			'A',
			'K',
			'Q',
			'T',
			'9',
			'8',
			'7',
			'6',
			'5',
			'4',
			'3',
			'2',
			'J',
		}
	}

	if a == b {
		return 0
	}
	var aIdx, bIdx int
	for i, v := range cmp {
		if v == a {
			aIdx = i
		}
		if v == b {
			bIdx = i
		}
	}
	if aIdx < bIdx {
		return 1
	}
	return -1
}

func getHands(input string, part int) []Hand {
	var hands []Hand
	lines := utils.ReadLines(input)

	for _, line := range lines {
		parts := strings.Split(line, " ")
		cards := parts[0]
		score := utils.ToInt(parts[1])

		handType := getHandType(cards)
		if part == 2 {
			handType = getHandTypePart2(cards)
		}

		hand := Hand{
			Cards: cards,
			Score: score,
			Hand:  handType,
		}

		hands = append(hands, hand)
	}

	return hands
}

func calcScore(hands []Hand) int {
	total := 0
	for i, hand := range hands {
		total += (i + 1) * hand.Score
	}
	return total
}

func Day07First(input string) int {
	part := 1
	hands := getHands(input, part)

	cmp := func(a, b Hand) int {
		if a.Hand < b.Hand {
			return -1
		}
		if a.Hand > b.Hand {
			return 1
		}

		for i := range a.Cards {
			compare := compareCards(rune(a.Cards[i]), rune(b.Cards[i]), part)
			if compare != 0 {
				return compare
			}
		}

		return 0
	}
	slices.SortFunc(hands, cmp)

	return calcScore(hands)
}

func Day07Second(input string) int {
	part := 2
	hands := getHands(input, part)

	cmp := func(a, b Hand) int {
		if a.Hand < b.Hand {
			return -1
		}
		if a.Hand > b.Hand {
			return 1
		}

		for i := range a.Cards {
			compare := compareCards(rune(a.Cards[i]), rune(b.Cards[i]), part)
			if compare != 0 {
				return compare
			}
		}

		return 0
	}
	slices.SortFunc(hands, cmp)

	return calcScore(hands)
}

func Day07() {
	fmt.Println(utils.DAY_PREFIX, "Day 07")
	fmt.Println(utils.PART_PREFIX, "Part 1:", Day07First(utils.ReadFile("y2023/input/day07.txt")))
	fmt.Println(utils.PART_PREFIX, "Part 2:", Day07Second(utils.ReadFile("y2023/input/day07.txt")))
}
