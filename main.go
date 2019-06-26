package main

import (
	"fmt"
)

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
	card1 := Card{points: "4", face: "H"}
	card2 := Card{points: "4", face: "D"}
	card3 := Card{points: "2", face: "H"}
	card4 := Card{points: "2", face: "S"}
	card5 := Card{points: "2", face: "C"}
	cards := [5]Card{card1, card2, card3, card4, card5}
	fmt.Println("kata poker")
	hands := PokerHands{cards}

	CountCardPoint := make(map[string]int)
	countCardFace := make(map[string]int)

	i := 0
	for i < len(hands.cards) {
		fmt.Println("card: ", hands.cards[i])
		card := hands.cards[i]

		_, poi := CountCardPoint[card.points]
		_, fac := countCardFace[card.face]
		if fac {
			countCardFace[card.face]++
		} else {
			countCardFace[card.face] = 1
		}
		if poi {
			CountCardPoint[card.points]++
		} else {
			CountCardPoint[card.points] = 1
		}
		i++
	}

	fmt.Println("CountCardPoint: ", CountCardPoint)
	fmt.Println("countCardFace: ", countCardFace)
	if len(countCardFace) == 1 {
		fmt.Println("Result: Flush")
	} else if len(CountCardPoint) == 2 {
		for _, num := range CountCardPoint {
			if num == 4 {
				fmt.Println("Result: Four of kind")
			} else if num == 3 {
				fmt.Println("Result: Full house")
			}
			// break
		}
	} else {
		fmt.Println("Result: High Card")
	}

}

func test(cards [5]Card) string {
	hands := PokerHands{cards}

	CountCardPoint := make(map[string]int)
	countCardFace := make(map[string]int)

	i := 0
	for i < len(hands.cards) {
		card := hands.cards[i]

		_, poi := CountCardPoint[card.points]
		_, fac := countCardFace[card.face]
		if fac {
			countCardFace[card.face]++
		} else {
			countCardFace[card.face] = 1
		}
		if poi {
			CountCardPoint[card.points]++
		} else {
			CountCardPoint[card.points] = 1
		}
		i++
	}

	// fmt.Println("CountCardPoint: ", CountCardPoint)
	// fmt.Println("countCardFace: ", countCardFace)
	if len(countCardFace) == 1 {
		fmt.Println("Result: Flush")
		return "Result: Flush"
	} else if len(CountCardPoint) == 2 {
		for _, num := range CountCardPoint {
			if num == 4 {
				return "Result: Four of kind"
			} else if num == 3 {
				return "Result: Full house"
			}
			// break
		}
	} else if len(CountCardPoint) == 3 {
		for _, num := range CountCardPoint {
			if num == 3 {
				return "Result: Three of kind"
			} else if num == 2 {
				return "Result: Two pair"
			}
			// break
		}
	} else if len(CountCardPoint) == 4 {
		for _, num := range CountCardPoint {
			if num == 2 {
				return "Result: One pair"
			}
			// break
		}
	} else if len(CountCardPoint) == 5 {
		return "Result: High card"
	}
	return ""
}
