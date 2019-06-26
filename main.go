package main

import "fmt"
// import "strconv"
import "github.com/bradfitz/slice"

var m_point = map[string]int{
	"2": 2, 
	"3": 3, 
	"4": 4, 
	"5": 5, 
	"6": 6, 
	"7": 7, 
	"8": 8, 
	"9": 9, 
	"10": 10, 
	"J": 11, 
	"Q": 12, 
	"K": 13, 
	"A": 14,
}

// Card -
type Card struct {
	points string // 2 - A
	face   string // S,H,D,C
}

// PokerHands -
type PokerHands struct {
	cards []Card
}

func (c Card) getPoint() int {
	return m_point[c.points];


	// return point["K"]

	// if c.points == "A" {
	// 	return 14
	// } else if c.points == "K" {
	// 	return 13
	// } else if c.points == "Q" {
	// 	return 12
	// } else if c.points == "J" {
	// 	return 11
	// }

	// s, _ := strconv.ParseInt(c.points, 10, 32)
	// return int(s)
}

func main() {
	
	card1 := Card{"10", "C"}
	card2 := Card{"10", "C"}
	card3 := Card{"9", "C"}
	card4 := Card{"8", "C"}
	card5 := Card{"7", "C"}
	cards := []Card{card1, card2, card3, card4, card5}

	pokerHands := PokerHands{cards}

	fmt.Println("Cards: ", cards);
	
	result := printHand(pokerHands)
	fmt.Print(result)
	
}

func printHand(hand PokerHands) string {
	slice.Sort(hand.cards[:], func(i, j int) bool {
		return hand.cards[i].getPoint() < hand.cards[j].getPoint()
	})

	result := ""

	var customDistinct = getDistinct(hand.cards)
	var cntDistinct = len(customDistinct)
	// fmt.Println("Count:", cntDistinct, " Data: ", customDistinct)
	
	if StraightFlush(hand.cards) {
		result = "Straight flush"
	} else if cntDistinct == 2 {
		if (customDistinct[0] == 4) {
			result = "Four of a kind"
		} else {
			result = "Full house"
		}	
	} else if Flush(hand.cards) {
		result = "Flush"
	} else if Straight(hand.cards) {
		result = "Straight"
	} else if cntDistinct == 3 {
		if customDistinct[0] == 3 {
			result = "Three of a kind"
		} else  {
			result = "Two pair"
		}
	} else if cntDistinct == 4 {
		result = "One pair"
	} else {
		result = "High card"
	}

	return result
}

func getDistinct(cards []Card) []int {
	var s []int

	var cnt = 1

	var strPoint = cards[0].points
	for i := 1; i < len(cards); i++ {
		if strPoint != cards[i].points {
			strPoint = cards[i].points
			s = append(s, cnt)

			cnt = 1
		} else {
			cnt++
		}
	}

	s = append(s, cnt)

	sortInt(s, false)
	return s
}

func sortInt(arr []int, sortAsc bool) {
	slice.Sort(arr[:], func(i, j int) bool {
		if sortAsc {
			return arr[i] < arr[j]
		} else {
			return arr[i] > arr[j]
		}
	})
}

func checkIsSameAllFaces(cards []Card) bool {
	for i := 1; i < len(cards); i++ {
		if cards[0].face != cards[i].face {
			return false
		}
	}

	return true
}

func checkIsStraight(cards []Card) bool {
	for i := 1; i < len(cards); i++ {
		if cards[i-1].getPoint() != (cards[i].getPoint() - 1) {
			return false
		}
	}

	return true
}

func StraightFlush(cards []Card) bool {

	isStraight := checkIsStraight(cards)
	isSameFace := checkIsSameAllFaces(cards)

	return isStraight && isSameFace
}

func Straight(cards []Card) bool {
	return checkIsStraight(cards)
}

func Flush(cards []Card) bool {
	return checkIsSameAllFaces(cards)
}
