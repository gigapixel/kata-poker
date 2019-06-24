package main

import "fmt"

// Card -
type Card struct {
	points string // 2 - A
	face   string // S,H,D,C
}

// PokerHands -
type PokerHands struct {
	cards [5]Card
}

func main() {
	// fmt.Println("kata poker")
	card1 := Card{"2", "S"}
	card2 := Card{"2", "H"}
	card3 := Card{"2", "D"}
	card4 := Card{"2", "C"}
	card5 := Card{"6", "S"}
	// hand := PokerHands(cards)
	// hand.print(card1)
	cards := [5]Card{card1, card2, card3, card4, card5}
	pokerHands := PokerHands{cards}

	result := printHand(pokerHands)
	fmt.Print(result)

}

func printHand(hand PokerHands) string {
	result := ""
	// pPoint := ""
	// cPoint := ""
	// pFace := ""
	countFace := 1

	// 4 cards
	for i := 0; i <= 3; i++ {
		if hand.cards[i].points == hand.cards[i+1].points {
			countFace++
		}
	}

	if countFace == 4 {
		result = "Four of a kind"
	}

	return result
}
