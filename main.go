package main

import "fmt"
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

func main() {
	// fmt.Println("kata poker")
	card1 := Card{"2", "S"}
	card2 := Card{"3", "S"}
	card3 := Card{"6", "S"}
	card4 := Card{"4", "S"}
	card5 := Card{"2", "S"}
	// hand := PokerHands(cards)
	// hand.print(card1)
	cards := [5]Card{card1, card2, card3, card4, card5}
	slice.Sort(cards[:], func(i, j int) bool {
		return cards[i].points < cards[j].points
	})
	// fmt.Print(cards)

	pokerHands := PokerHands{cards}

	result := printHand(pokerHands)
	fmt.Print(result)

}

func printHand(hand PokerHands) string {
	result := ""
	// pPoint := ""
	countPoint := 1
	// pFace := ""
	countFace := 1

	// Straight flush

	// Four of a kind
	for i := 0; i <= 3; i++ {
		if hand.cards[i].points == hand.cards[i+1].points {
			countPoint++
		}
	}

	// Full house

	// flush
	for i := 0; i <= 3; i++ {
		if hand.cards[i].face == hand.cards[i+1].face {
			countFace++
		}
	}

	// Straight
	// Three of a kind
	// Two pair
	// One pair

	if countPoint == 4 {
		result = "Four of a kind"
	} else if countFace == 5 {
		result = "Flush"
	} else if countPoint == 3 {
		result = "Three of a kind"
	} else if countPoint == 2 {
		result = "One pair"
	}
	return result
}
