package main

import (
	"sort"
	"strconv"
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
	switch points {
	case "J":
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 14
	default:
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
	for i := 1; i < len(cards); i++ {
		switch (cards[i-1].value + 1) != cards[i].value {
		case true:
			return false
		}
	}
	return true
}

func isFlush(faces map[string]int) bool {
	return len(faces) == 1
}

func poker(cards []Card) interface{} {
	assignPointsValue(cards)
	sort.Sort(ByValue(cards))

	faces := groupAndCountByFace(cards)
	points := groupAndCountByPoints(cards)

	straight := isStraight(cards)
	flush := isFlush(faces)
	// fourOfaKind := false
	threeOfaKind := false
	firstPair := false
	secondPair := false

	for _, count := range points {
		switch count {
		case 4:
			return "four of a kind"
		case 3:
			threeOfaKind = true
		case 2:
			switch !firstPair {
			case true:
				firstPair = true
			default:
				switch !secondPair {
				case true:
					secondPair = true
				}
			}
		}
	}

	switch straight && flush {
	case true:
		return "straight flush"
	}

	switch straight {
	case true:
		return "straight"
	}

	switch flush {
	case true:
		return "flush"
	}

	twoPair := firstPair && secondPair
	onePair := !twoPair && (firstPair || secondPair)

	switch twoPair {
	case true:
		return "two pair"
	}

	switch threeOfaKind && onePair {
	case true:
		return "full house"
	}

	switch threeOfaKind {
	case true:
		return "three of a kind"
	}

	switch onePair {
	case true:
		return "one pair"
	}

	return "high card"
}