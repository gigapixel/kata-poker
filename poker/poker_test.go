package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestStraightFlush(t *testing.T) {
	cards := []Card{
		Card{points: "J", face: "C"},
		Card{points: "7", face: "C"},
		Card{points: "9", face: "C"},
		Card{points: "10", face: "C"},
		Card{points: "8", face: "C"},
	}
	actual := poker(cards)
	assert.Equal(t, "straight flush", actual, "poker cards on hand should be straight flush")
}

func TestFourOfKind(t *testing.T) {
	cards := []Card{
		Card{points: "5", face: "C"},
		Card{points: "5", face: "D"},
		Card{points: "5", face: "H"},
		Card{points: "5", face: "S"},
		Card{points: "2", face: "D"},
	}
	actual := poker(cards)
	assert.Equal(t, "four of a kind", actual, "poker cards on hand should be four of kind")
}

func TestFullHouse(t *testing.T) {
	cards := []Card{
		Card{points: "6", face: "S"},
		Card{points: "6", face: "H"},
		Card{points: "6", face: "D"},
		Card{points: "K", face: "C"},
		Card{points: "K", face: "H"},
	}
	actual := poker(cards)
	assert.Equal(t, "full house", actual, "poker cards on hand should be full house")
}

func TestFlush(t *testing.T) {
	cards := []Card{
		Card{points: "J", face: "D"},
		Card{points: "9", face: "D"},
		Card{points: "8", face: "D"},
		Card{points: "4", face: "D"},
		Card{points: "3", face: "D"},
	}
	actual := poker(cards)
	assert.Equal(t, "flush", actual, "poker cards on hand should be flush")
}

func TestStraight(t *testing.T) {
	cards := []Card{
		Card{points: "10", face: "D"},
		Card{points: "9", face: "S"},
		Card{points: "8", face: "H"},
		Card{points: "7", face: "D"},
		Card{points: "6", face: "C"},
	}
	actual := poker(cards)
	assert.Equal(t, "straight", actual, "poker cards on hand should be straight")
}

func TestThreeOfKind(t *testing.T) {
	cards := []Card{
		Card{points: "Q", face: "C"},
		Card{points: "Q", face: "S"},
		Card{points: "Q", face: "H"},
		Card{points: "9", face: "H"},
		Card{points: "2", face: "S"},
	}
	actual := poker(cards)
	assert.Equal(t, "three of a kind", actual, "poker cards on hand should be three of kind")
}

func TestTwoPair(t *testing.T) {
	cards := []Card{
		Card{points: "J", face: "H"},
		Card{points: "J", face: "S"},
		Card{points: "3", face: "C"},
		Card{points: "3", face: "S"},
		Card{points: "2", face: "H"},
	}
	actual := poker(cards)
	assert.Equal(t, "two pair", actual, "poker cards on hand should be two pair")
}

func TestOnePair(t *testing.T) {
	cards := []Card{
		Card{points: "10", face: "S"},
		Card{points: "10", face: "H"},
		Card{points: "8", face: "S"},
		Card{points: "7", face: "H"},
		Card{points: "4", face: "C"},
	}
	actual := poker(cards)
	assert.Equal(t, "one pair", actual, "poker cards on hand should be one pair")
}

func TestHighCard(t *testing.T) {
	cards := []Card{
		Card{points: "K", face: "D"},
		Card{points: "Q", face: "D"},
		Card{points: "7", face: "S"},
		Card{points: "3", face: "S"},
		Card{points: "A", face: "H"},
	}
	actual := poker(cards)
	assert.Equal(t, "high card", actual, "poker cards on hand should be high card")
}
