package main

import "fmt"
import "strconv"
import "github.com/bradfitz/slice"

// Card -
type Card struct {
	points string // 2 - A
	face   string // S,H,D,C
}

// PokerHands -
type PokerHands struct {
	cards [5]Card
}

func (c Card) getPoint() int {
	if c.points == "A" {
		return 14
	} else if c.points == "K" {
		return 13
	} else if c.points == "Q" {
		return 12
	} else if c.points == "J" {
		return 11
	}

	s, _ := strconv.ParseInt(c.points, 10, 32)
	return int(s)
}

func main() {
	// fmt.Println("kata poker")
	card1 := Card{"10", "C"}
	card2 := Card{"J", "C"}
	card3 := Card{"9", "C"}
	card4 := Card{"8", "C"}
	card5 := Card{"7", "C"}
	// hand := PokerHands(cards)
	// hand.print(card1)
	cards := [5]Card{card1, card2, card3, card4, card5}

	

	pokerHands := PokerHands{cards}

	result := printHand(pokerHands)
	fmt.Print(result)
	//return result
}

func printHand(hand PokerHands) string {
	slice.Sort(hand.cards[:], func(i, j int) bool {
		return hand.cards[i].getPoint() < hand.cards[j].getPoint()
	})

	result := ""	

	if (StraightFlush(hand.cards)) {
		result = "Straight flush"
	} else if (FourOfAKind(hand.cards)) {
		result = "Four of a kind"
	} else if (FullHouse(hand.cards)) {
		result = "Full house"
	} else if (Flush(hand.cards)) {
		result = "Flush"
	} else if (Straight(hand.cards)) {
		result = "Straight"
	}	else if (ThreeOfAKind(hand.cards)) {
		result = "Three of a kind"
	} else if (TwoPair(hand.cards)) {
		result = "Two pair"
	} else if (OnePair(hand.cards)) {
		result = "One pair"
	} else {
		result = "High card"
	}

	return result
}

func sameAllFaces(cards [5]Card) bool {
	for i := 1; i < len(cards); i++ {
		if cards[0].face != cards[i].face {
			return false
		}
	}

	return true
}


func StraightFlush(cards [5]Card) bool {
	pt0 := cards[0].getPoint()
	pt1 := cards[1].getPoint()
	pt2 := cards[2].getPoint()
	pt3 := cards[3].getPoint()
	pt4 := cards[4].getPoint()

	isCorrectPoint := pt0+1 == pt1 && pt0+2 == pt2 && pt0+3 == pt3 && pt0+4 == pt4
	return isCorrectPoint && sameAllFaces(cards)
}

func FourOfAKind(cards [5]Card) bool {
	return false;
}

func FullHouse(cards [5]Card) bool {
	return false;
}

func Straight(cards [5]Card) bool {
	pt0 := cards[0].getPoint()
	pt1 := cards[1].getPoint()
	pt2 := cards[2].getPoint()
	pt3 := cards[3].getPoint()
	pt4 := cards[4].getPoint()

	isCorrectPoint := pt0+1 == pt1 && pt0+2 == pt2 && pt0+3 == pt3 && pt0+4 == pt4
	return isCorrectPoint
}

func Flush(cards [5]Card) bool {
	return sameAllFaces(cards);
}

func ThreeOfAKind(cards [5]Card) bool {
	return false;
}

func TwoPair(cards [5]Card) bool {
	return false;
}

func OnePair(cards [5]Card) bool {
	return false;
}

// func FourOfKind(cards [5]Card) bool {
// 	var cntDup := 0;
// 	for i := 0; i < len(cards); i++ {

// 	}

// 	return false
// }
