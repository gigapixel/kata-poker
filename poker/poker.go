package main

import (
	"sort"
	"strconv"
	"strings"
)

// Card -
type Card struct {
	points string // 2 - A
	face   string // S(Spade), H(Heart), D(Diamond), C(Club)
	value  int
}

// PokerHands -
type PokerHands struct {
	cards [5]Card
}

type ByValue []Card

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i].value < a[j].value }

func getPointsValue(points string) int {
	if strings.EqualFold("J", points) {
		return 11
	} else if strings.EqualFold("Q", points) {
		return 12
	} else if strings.EqualFold("K", points) {
		return 13
	} else if strings.EqualFold("A", points) {
		return 14
	} else {
		num, _ := strconv.Atoi(points)
		return num
	}
}

func assignPointsValue(cards []Card) {
	for i := 0; i < len(cards); i++ {
		cards[i].value = getPointsValue(cards[i].points)
	}
}

func groupAndCountByFace(cards []Card) map[string]int {
	m := map[string]int{}
	for _, c := range cards {
		m[c.face]++
	}
	return m
}

func groupAndCountByPoints(cards []Card) map[string]int {
	m := map[string]int{}
	for _, c := range cards {
		m[c.points]++
	}
	return m
}

func isStraight(cards []Card) bool {
	for i, card := range cards {
		if i == 0 {
			continue
		}
		if (cards[i-1].value + 1) != card.value {
			return false
		}
	}
	return true
}

func isFlush(faces map[string]int) bool {
	return len(faces) == 1
}

// func printCardsTable(cards []Card) {
// 	for _, c := range cards {
// 		fmt.Println("Card[", c.face, "][", c.points, "] = ", c.value)
// 	}
// }

func poker(cards []Card) interface{} {
	assignPointsValue(cards)
	sort.Sort(ByValue(cards))

	faces := groupAndCountByFace(cards)
	points := groupAndCountByPoints(cards)

	straight := isStraight(cards)
	flush := isFlush(faces)
	fourOfaKind := false
	threeOfaKind := false
	firstPair := false
	secondPair := false

	for _, count := range points {
		if count == 4 {
			fourOfaKind = true
		}

		if count == 3 {
			threeOfaKind = true
		}

		if count == 2 && !firstPair {
			firstPair = true
		} else if count == 2 && !secondPair {
			secondPair = true
		}
	}

	if straight && flush {
		return "straight flush"
	}

	if straight {
		return "straight"
	}

	if flush {
		return "flush"
	}

	if fourOfaKind {
		return "four of a kind"
	}

	twoPair := firstPair && secondPair
	onePair := !twoPair && (firstPair || secondPair)

	if twoPair {
		return "two pair"
	}

	if threeOfaKind && onePair {
		return "full house"
	}

	if threeOfaKind {
		return "three of a kind"
	}

	if onePair {
		return "one pair"
	}

	return "high card"
}
