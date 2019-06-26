package main

import "testing"
import "github.com/stretchr/testify/assert"

type Person struct {
	name string
}

// func TestFlush(t *testing.T) {
// 	cards := [5]Card{Card{"A", "S"}, Card{"2", "S"}, Card{"3", "S"}, Card{"4", "S"}, Card{"5", "S"}}
// 	hand := PokerHands{cards}
// 	var res = cal(hand);
// 	var expected = "flush";
// 	assert.Equal(t, expected, res, "Result should be 30")
// }

func TestFiveOfAkind(t *testing.T) {
	cards := [5]Card{
		{"A", "S"}, 
		{"A", "H"}, 
		{"A", "D"}, 
		{"A", "C"}, 
		{"J", "J"},
	}
	hand := PokerHands{cards}
	var res = cal(hand);
	var expected = "Five of a kind";
	assert.Equal(t, expected, res, "Result should be 30")
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
	var res = cal(hand);
	var expected = "StraightFlush";
	assert.Equal(t, expected, res, "Result should be 30")
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
	var res = cal(hand);
	var expected = "Flush";
	assert.Equal(t, expected, res, "Result should be 30")
}

func TestFullHouse(t *testing.T) {
	cards := [5]Card{
		Card{"6", "S"}, 
		Card{"6", "D"}, 
		Card{"6", "H"}, 
		Card{"11", "C"}, 
		Card{"11", "K"},
	}
	hand := PokerHands{cards}
	var res = cal(hand);
	var expected = "Full house";
	assert.Equal(t, expected, res, "Result should be 30")
}

func TestThreeOfAKind(t *testing.T) {
	cards := [5]Card{
		Card{"6", "S"}, 
		Card{"6", "D"}, 
		Card{"6", "H"}, 
		Card{"9", "C"}, 
		Card{"11", "K"},
	}
	hand := PokerHands{cards}
	var res = cal(hand);
	var expected = "Three of a kind";
	assert.Equal(t, expected, res, "Result should be 30")
}

func TestTwoPair(t *testing.T) {
	cards := [5]Card{
		Card{"6", "S"}, 
		Card{"6", "D"}, 
		Card{"9", "H"}, 
		Card{"9", "C"}, 
		Card{"11", "K"},
	}
	hand := PokerHands{cards}
	var res = cal(hand);
	var expected = "Two pair";
	assert.Equal(t, expected, res, "Result should be 30")
}

func TestOnePair(t *testing.T) {
	cards := [5]Card{
		Card{"6", "S"}, 
		Card{"6", "D"}, 
		Card{"8", "H"}, 
		Card{"9", "C"}, 
		Card{"11", "K"},
	}
	hand := PokerHands{cards}
	var res = cal(hand);
	var expected = "One pair";
	assert.Equal(t, expected, res, "Result should be 30")
}

func TestHighCard(t *testing.T) {
	cards := [5]Card{
		Card{"3", "S"}, 
		Card{"6", "D"}, 
		Card{"8", "H"}, 
		Card{"9", "C"}, 
		Card{"11", "K"},
	}
	hand := PokerHands{cards}
	var res = cal(hand);
	var expected = "High card";
	assert.Equal(t, expected, res, "Result should be 30")
}


