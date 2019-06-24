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
	card1 := Card{points: "A", face: "S"}
	card2 := Card{points: "2", face: "S"}
	card3 := Card{points: "5", face: "S"}
	card4 := Card{points: "8", face: "S"}
	card5 := Card{points: "4", face: "S"}
	cards := [5]Card{card1, card2, card3, card4, card5}
	fmt.Println("kata poker")
	hands := PokerHands{cards}

	S := 0
	D := 0
	H := 0
	C := 0

	i := 0
	for i < len(hands.cards) {
		fmt.Println("card: ", hands.cards[i])
		card := hands.cards[i]

		// xxx := [...]string{}
		// xxx[] = card.point

		if card.face == "S" {
			S++
		}
		if card.face == "D" {
			D++
		}
		if card.face == "H" {
			H++
		}
		if card.face == "C" {
			C++
		}
		i++
	}

	if S == 5 || D == 5 || H == 5 || C == 5 {
		fmt.Println("Result: Flush")
	} else {
		fmt.Println("Result: High Card")
	}

}
