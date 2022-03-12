// package yacht implements scoring of a yacht game
package yacht

import (
	"sort"
)

// Sum returns the total of dice with value v
func Sum(dice []int, v int) int {
	total := 0
	for _, d := range dice {
		if d == v {
			total += d
		}
	}
	return total
}

// FullHouse returns the total dice values of a full house
func FullHouse(dice []int) int {
	sort.Ints(dice)
	if dice[0] == dice[4] {
		return 0
	}
	if dice[0] == dice[1] && dice[3] == dice[4] {
		if dice[1] == dice[2] || dice[2] == dice[3] {
			return Choice(dice)
		}
	}
	return 0
}

// FourOfKind returns the total dice values of four of a kind
func FourOfKind(dice []int) int {
	sort.Ints(dice)
	if dice[1] == dice[2] && dice[2] == dice[3] {
		if dice[0] == dice[1] || dice[3] == dice[4] {
			return dice[1] * 4
		}
	}
	return 0
}

// Straight returns 30 if the dice is a straight
func Straight(dice []int, first int, last int) int {
	sort.Ints(dice)
	if dice[0] < dice[1] && dice[1] < dice[2] && dice[2] < dice[3] && dice[3] < dice[4] {
		if dice[0] == first && dice[4] == last {
			return 30
		}
	}
	return 0
}

// Choice returns the total dice values
func Choice(dice []int) int {
	total := 0
	for _, d := range dice {
		total += d
	}
	return total
}

// Yacht returns 50 if all dice values are the same value
func Yacht(dice []int) int {
	if dice[0] == dice[1] && dice[1] == dice[2] && dice[2] == dice[3] && dice[3] == dice[4] {
		return 50
	}
	return 0
}

// Score returns the score of a dice slice given a category
func Score(dice []int, category string) int {
	score := 0
	switch category {
	case "ones":
		score = Sum(dice, 1)
	case "twos":
		score = Sum(dice, 2)
	case "threes":
		score = Sum(dice, 3)
	case "fours":
		score = Sum(dice, 4)
	case "fives":
		score = Sum(dice, 5)
	case "sixes":
		score = Sum(dice, 6)
	case "full house":
		score = FullHouse(dice)
	case "four of a kind":
		score = FourOfKind(dice)
	case "little straight":
		score = Straight(dice, 1, 5)
	case "big straight":
		score = Straight(dice, 2, 6)
	case "choice":
		score = Choice(dice)
	case "yacht":
		score = Yacht(dice)
	}
	return score
}
