package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestFourOfAKind(t *testing.T) {
	card1 := Card{"1", "S"}
	card2 := Card{"1", "H"}
	card3 := Card{"1", "D"}
	card4 := Card{"1", "C"}
	card5 := Card{"5", "S"}
	hand := [5]Card{card1, card2, card3, card4, card5}
	pokerHand := PokerHands{hand}

	var result = printHand(pokerHand)

	assert.Equal(t, "Four of a kind", result, "Four of a kind")
}

func TestStraightFlush(t *testing.T) {
	card1 := Card{"10", "C"}
	card2 := Card{"J", "C"}
	card3 := Card{"9", "C"}
	card4 := Card{"8", "C"}
	card5 := Card{"7", "C"}
	hand := [5]Card{card1, card2, card3, card4, card5}
	pokerHand := PokerHands{hand}

	var result = printHand(pokerHand)

	assert.Equal(t, "Straight flush", result, "Straight flush")
}
