package main

import "testing"
import "github.com/stretchr/testify/assert"

type Person struct {
	name string
}

func TestFourOfKind(t *testing.T) {
	cards := [5]Card{
	 Card{"5", "S"},
	 Card{"5", "H"},
	 Card{"5", "D"},
	 Card{"5", "C"},
	 Card{"8", "H"},
	}
	hand := PokerHands{cards}
	var res = cal(hand)
	var expected = "Four of a kind"
	assert.Equal(t, expected, res, "Result should be four of a kind")
 }

func TestFiveOfAkind(t *testing.T) {
	cards := [5]Card{
		{"A", "S"},
		{"A", "H"},
		{"A", "D"},
		{"A", "C"},
		{"J", "J"},
	}
	hand := PokerHands{cards}
	var res = cal(hand)
	var expected = "Five of a kind"
	assert.Equal(t, expected, res, "Result should be five of kind")
}

func TestStraightFlush(t *testing.T) {
	cards := [5]Card{
		Card{"J", "S"},
		Card{"10", "S"},
		Card{"9", "S"},
		Card{"8", "S"},
		Card{"7", "S"},
	}
	hand := PokerHands{cards}
	var res = cal(hand)
	var expected = "StraightFlush"
	assert.Equal(t, expected, res, "Result should be straight flush")
}

func TestFlush(t *testing.T) {
	cards := [5]Card{
		Card{"J", "D"},
		Card{"9", "D"},
		Card{"8", "D"},
		Card{"4", "D"},
		Card{"3", "D"},
	}
	hand := PokerHands{cards}
	var res = cal(hand)
	var expected = "Flush"
	assert.Equal(t, expected, res, "Result should be flush")
}

func TestFullHouse(t *testing.T) {
	cards := [5]Card{
		Card{"6", "S"},
		Card{"6", "D"},
		Card{"6", "H"},
		Card{"J", "C"},
		Card{"J", "H"},
	}
	hand := PokerHands{cards}
	var res = cal(hand)
	var expected = "Full house"
	assert.Equal(t, expected, res, "Result should be full house")
}

func TestThreeOfAKind(t *testing.T) {
	cards := [5]Card{
		Card{"6", "S"},
		Card{"6", "D"},
		Card{"6", "H"},
		Card{"9", "C"},
		Card{"J", "H"},
	}
	hand := PokerHands{cards}
	var res = cal(hand)
	var expected = "Three of a kind"
	assert.Equal(t, expected, res, "Result should be three of a kind")
}

func TestTwoPair(t *testing.T) {
	cards := [5]Card{
		Card{"6", "S"},
		Card{"6", "D"},
		Card{"9", "H"},
		Card{"9", "C"},
		Card{"J", "C"},
	}
	hand := PokerHands{cards}
	var res = cal(hand)
	var expected = "Two pair"
	assert.Equal(t, expected, res, "Result should be two pair")
}

func TestOnePair(t *testing.T) {
	cards := [5]Card{
		Card{"6", "S"},
		Card{"6", "D"},
		Card{"8", "H"},
		Card{"9", "C"},
		Card{"J", "C"},
	}
	hand := PokerHands{cards}
	var res = cal(hand)
	var expected = "One pair"
	assert.Equal(t, expected, res, "Result should be one pair")
}

func TestHighCard(t *testing.T) {
	cards := [5]Card{
		Card{"3", "S"},
		Card{"6", "D"},
		Card{"8", "H"},
		Card{"9", "C"},
		Card{"J", "C"},
	}
	hand := PokerHands{cards}
	var res = cal(hand)
	var expected = "High card"
	assert.Equal(t, expected, res, "Result should be high card")
}
