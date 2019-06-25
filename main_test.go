package main

import "testing"
import "fmt"
import "github.com/stretchr/testify/assert"

func TestStraightFlush(t *testing.T) {
	card1 := Card{"10", "C"}
	card2 := Card{"J", "C"}
	card3 := Card{"9", "C"}
	card4 := Card{"8", "C"}
	card5 := Card{"7", "C"}
	hand := [5]Card{card1, card2, card3, card4, card5}
	pokerHand := PokerHands{hand}

	var result = printHand(pokerHand)
	fmt.Println(result);
	

	assert.Equal(t, "Straight flush", result, "Straight flush")
}
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

func TestFullHouse(t *testing.T) {
	card1 := Card{"6", "S"}
	card2 := Card{"6", "H"}
	card3 := Card{"6", "D"}
	card4 := Card{"K", "S"}
	card5 := Card{"K", "C"}
	hand := [5]Card{card1, card2, card3, card4, card5}
	pokerHand := PokerHands{hand}

	var result = printHand(pokerHand)

	assert.Equal(t, "Full house", result, "Full house")
}

func TestFlush(t *testing.T) {
	card1 := Card{"7", "S"}
	card2 := Card{"6", "S"}
	card3 := Card{"3", "S"}
	card4 := Card{"4", "S"}
	card5 := Card{"K", "S"}
	hand := [5]Card{card1, card2, card3, card4, card5}
	pokerHand := PokerHands{hand}

	var result = printHand(pokerHand)

	assert.Equal(t, "Flush", result, "Flush")
}

func TestStraight(t *testing.T) {
	card1 := Card{"10", "S"}
	card2 := Card{"9", "H"}
	card3 := Card{"8", "D"}
	card4 := Card{"7", "C"}
	card5 := Card{"6", "S"}
	hand := [5]Card{card1, card2, card3, card4, card5}
	pokerHand := PokerHands{hand}

	var result = printHand(pokerHand)

	assert.Equal(t, "Straight", result, "Straight")
}

func TestThreeOfAKind(t *testing.T) {
	card1 := Card{"Q", "S"}
	card2 := Card{"Q", "H"}
	card3 := Card{"Q", "D"}
	card4 := Card{"7", "C"}
	card5 := Card{"6", "S"}
	hand := [5]Card{card1, card2, card3, card4, card5}
	pokerHand := PokerHands{hand}

	var result = printHand(pokerHand)

	assert.Equal(t, "Three of a kind", result, "Three of a kind")
}

func TestTwoPair(t *testing.T) {
	card1 := Card{"Q", "S"}
	card2 := Card{"Q", "H"}
	card3 := Card{"3", "D"}
	card4 := Card{"3", "C"}
	card5 := Card{"6", "S"}
	hand := [5]Card{card1, card2, card3, card4, card5}
	pokerHand := PokerHands{hand}

	var result = printHand(pokerHand)

	assert.Equal(t, "Two pair", result, "Two pair")
}

func TestOnePair(t *testing.T) {
	card1 := Card{"Q", "S"}
	card2 := Card{"Q", "H"}
	card3 := Card{"3", "D"}
	card4 := Card{"5", "C"}
	card5 := Card{"6", "S"}
	hand := [5]Card{card1, card2, card3, card4, card5}
	pokerHand := PokerHands{hand}

	var result = printHand(pokerHand)

	assert.Equal(t, "One pair", result, "One pair")
}

func TestHighCard(t *testing.T) {
	card1 := Card{"Q", "S"}
	card2 := Card{"3", "H"}
	card3 := Card{"A", "D"}
	card4 := Card{"5", "C"}
	card5 := Card{"6", "S"}
	hand := [5]Card{card1, card2, card3, card4, card5}
	pokerHand := PokerHands{hand}

	var result = printHand(pokerHand)

	assert.Equal(t, "High card", result, "High card")
}
