package main

import (
	"testing"
)

func TestAlwaysReturn1(t *testing.T) {
	card1 := Card{points: "A", face: "S"}
	card2 := Card{points: "2", face: "S"}
	card3 := Card{points: "5", face: "S"}
	card4 := Card{points: "8", face: "S"}
	card5 := Card{points: "4", face: "S"}
	cards := [5]Card{card1, card2, card3, card4, card5}
	result := test(cards)
	if result != string("Result: Flush") {
		t.Error("Result should be Result: Flush")
	}
}

func TestAlwaysReturn2(t *testing.T) {
	card1 := Card{points: "A", face: "S"}
	card2 := Card{points: "A", face: "D"}
	card3 := Card{points: "A", face: "C"}
	card4 := Card{points: "A", face: "H"}
	card5 := Card{points: "4", face: "S"}
	cards := [5]Card{card1, card2, card3, card4, card5}
	result := test(cards)
	if result != string("Result: Four of kind") {
		t.Error("Result should be Result: Four of kind")
	}
}

func TestAlwaysReturn3(t *testing.T) {
	card1 := Card{points: "A", face: "S"}
	card2 := Card{points: "A", face: "D"}
	card3 := Card{points: "A", face: "C"}
	card4 := Card{points: "4", face: "H"}
	card5 := Card{points: "4", face: "S"}
	cards := [5]Card{card1, card2, card3, card4, card5}
	result := test(cards)
	if result != string("Result: Full house") {
		t.Error("Result should be Result: Full house")
	}
}

func TestAlwaysReturn4(t *testing.T) {
	card1 := Card{points: "A", face: "S"}
	card2 := Card{points: "A", face: "D"}
	card3 := Card{points: "A", face: "C"}
	card4 := Card{points: "4", face: "H"}
	card5 := Card{points: "5", face: "S"}
	cards := [5]Card{card1, card2, card3, card4, card5}
	result := test(cards)
	if result != string("Result: Three of kind") {
		t.Error("Result should be Result: Three of kind")
	}
}

func TestAlwaysReturn5(t *testing.T) {
	card1 := Card{points: "A", face: "S"}
	card2 := Card{points: "A", face: "D"}
	card3 := Card{points: "2", face: "C"}
	card4 := Card{points: "4", face: "H"}
	card5 := Card{points: "4", face: "S"}
	cards := [5]Card{card1, card2, card3, card4, card5}
	result := test(cards)
	if result != string("Result: Two pair") {
		t.Error("Result should be Result: Two pair")
	}
}

func TestAlwaysReturn6(t *testing.T) {
	card1 := Card{points: "A", face: "S"}
	card2 := Card{points: "A", face: "D"}
	card3 := Card{points: "2", face: "C"}
	card4 := Card{points: "3", face: "H"}
	card5 := Card{points: "4", face: "S"}
	cards := [5]Card{card1, card2, card3, card4, card5}
	result := test(cards)
	if result != string("Result: One pair") {
		t.Error("Result should be Result: One pair")
	}
}

func TestAlwaysReturn7(t *testing.T) {
	card1 := Card{points: "A", face: "S"}
	card2 := Card{points: "K", face: "D"}
	card3 := Card{points: "2", face: "C"}
	card4 := Card{points: "3", face: "H"}
	card5 := Card{points: "4", face: "S"}
	cards := [5]Card{card1, card2, card3, card4, card5}
	result := test(cards)
	if result != string("Result: High card") {
		t.Error("Result should be Result: High card")
	}
}
