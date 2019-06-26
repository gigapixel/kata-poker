package main 

import "testing" 
// import "github.com/stretchr/testify/assert" 


func TestAlwaysReturnFlush(t *testing.T) { 
input := [5]Card{Card{"8", "H"}, Card{"2", "H"}, Card{"3", "H"}, Card{"4", "H"}, Card{"5", "H"}} 
hands := PokerHands{input} 
result := getOutput(hands) 
if result != "Flush" { 
t.Error("Result should be Flush") 
} 
} 

func TestAlwaysReturnFourOfAKind(t *testing.T) { 
input := [5]Card{Card{"5", "H"}, Card{"5", "D"}, Card{"5", "S"}, Card{"5", "A"}, Card{"8", "H"}} 
hands := PokerHands{input} 
result := getOutput(hands) 
if result != "Four of a kind" { 
t.Error("Result should be Four of a kind") 
} 
} 

func TestAlwaysReturnFullHouse(t *testing.T) { 
input := [5]Card{Card{"5", "H"}, Card{"5", "D"}, Card{"5", "S"}, Card{"K", "D"}, Card{"K", "H"}} 
hands := PokerHands{input} 
result := getOutput(hands) 
if result != "Full house" { 
t.Error("Result should be Full house") 
} 
} 


// new 
func TestAlwaysReturnStraightFlush(t *testing.T) { 
input := [5]Card{Card{"J", "H"}, Card{"10", "H"}, Card{"9", "H"}, Card{"8", "H"}, Card{"7", "H"}} 
hands := PokerHands{input} 
result := getOutput(hands) 
if result != "Straight flush" { 
t.Error("Result should be Straight flush") 
} 
} 

func TestAlwaysReturnStraight(t *testing.T) { 
input := [5]Card{Card{"10", "H"}, Card{"9", "D"}, Card{"8", "S"}, Card{"7", "H"}, Card{"6", "C"}} 
hands := PokerHands{input} 
result := getOutput(hands) 
if result != "Straight" { 
t.Error("Result should be Straight") 
} 
} 

func TestAlwaysReturnThreeOfAKind(t *testing.T) { 
input := [5]Card{Card{"Q", "H"}, Card{"Q", "D"}, Card{"9", "H"}, Card{"2", "C"},Card{"Q", "S"}} 
hands := PokerHands{input} 
result := getOutput(hands) 
if result != "Three of a kind" { 
t.Error("Result should be Three of a kind") 
} 
} 

func TestAlwaysReturnTwoPair(t *testing.T) { 
input := [5]Card{Card{"J", "H"}, Card{"J", "D"}, Card{"3", "S"}, Card{"3", "H"}, Card{"2", "C"}} 
hands := PokerHands{input} 
result := getOutput(hands) 
if result != "Two pair" { 
t.Error("Result should be Two pair") 
} 
} 

func TestAlwaysReturnOnePair(t *testing.T) { 
input := [5]Card{Card{"10", "H"}, Card{"10", "D"}, Card{"8", "S"}, Card{"7", "H"}, Card{"4", "C"}} 
hands := PokerHands{input} 
result := getOutput(hands) 
if result != "One pair" { 
t.Error("Result should be One pair") 
} 
} 

func TestAlwaysReturnHighCard(t *testing.T) { 
input := [5]Card{Card{"K", "H"}, Card{"Q", "H"}, Card{"7", "S"}, Card{"4", "H"}, Card{"3", "C"}} 
hands := PokerHands{input} 
result := getOutput(hands) 
if result != "High card" { 
t.Error("Result should be High card") 
} 
}