package main

import "fmt"
import "sort"

// Card -
type Card struct {
	points string // 2 - A
	face   string // S,H,D,C
}

// PokerHands -
type PokerHands struct {
	cards [5]Card
}

func poker(hand PokerHands) {
	// fmt.Println("poker card on hand ->", hand)
	isStraight := false
	isFlush := false
	for i := 0; i < len(hand.cards); i++ {
		fmt.Println("card[", i+1, "] = ", hand.cards[i])
	}
}

func (c PokerHands) Less(i, j int) bool {
	fmt.Println(c.cards)
	return c.cards[i].points < c.cards[i].points
}

func main() {
	fmt.Println("kata poker")

	// validCardNumbers = ["2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"];
	// cardValues = [2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14];
	// validCardSuits = ["H", "D", "C", "S"];
	// var J=11, Q=12, K=13, A=14, C=1, D=2, H=4, S=8;

	cards1 := [5]Card{Card{"5", "C"}, Card{"7", "C"}, Card{"2", "C"}, Card{"3", "C"}, Card{"9", "C"}}
	hands1 := PokerHands{cards: cards1}
	sort.Sort(hands1)
	fmt.Println(hands1)
}
