package main

import "fmt"
import "sort"
import "strconv"

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
	fmt.Println("kata poker")

	c1 := Card{points: "4", face: "S"}
	c2 := Card{points: "Q", face: "S"}
	c3 := Card{points: "J", face: "S"}
	c4 := Card{points: "K", face: "S"}
	c5 := Card{points: "5", face: "S"}

	pokerHands := PokerHands{};
	pokerHands.cards[0] = c1;
	pokerHands.cards[1] = c2;
	pokerHands.cards[2] = c3;
	pokerHands.cards[3] = c4;
	pokerHands.cards[4] = c5;
	fmt.Println(pokerHands)
	
	r1 := SameFace(pokerHands);
	fmt.Println("Same Face : ", r1);

	if(IsFlsh(pokerHands)) {
		fmt.Println("Flsh");
	}

	IsContinueNumber(pokerHands, 3);
}

func IsContinueNumber(hands PokerHands, n) {
	var score []int
	for i := 0; i < 5; i++ {
		if(hands.cards[i].points == "K") {
			score = append(score, 15);
		} else if(hands.cards[i].points == "Q") {
			score = append(score, 14);
		} else if(hands.cards[i].points == "J") {
			score = append(score, 13);
		} else {
			k, _ := strconv.Atoi(hands.cards[i].points)
    	score = append(score, k);
		}
	}

	sort.Ints(score)

	if(score[0] )

	for i := 0; i < 5; i++ {
		fmt.Println("score : ", score[i]);
	}
}

func SameNumber(hands PokerHands) int {
	var n2 = 0;
	var n3 = 0;
	var n4 = 0;
	var C = 0;
	
	for i := 0; i < 5; i++ {
		if(hands.cards[i].face == "S") {
			S++;
		} else if(hands.cards[i].face == "H") {
			H++;
		} else if(hands.cards[i].face == "D") {
			D++;
		} else if(hands.cards[i].face == "C") {
			C++;
		}
	}

	ints := []int{S, H, D, C}

	sort.Ints(ints)
	return ints[3];
}

func SameFace(hands PokerHands) int {
	var S = 0;
	var H = 0;
	var D = 0;
	var C = 0;
	
	for i := 0; i < 5; i++ {
		if(hands.cards[i].face == "S") {
			S++;
		} else if(hands.cards[i].face == "H") {
			H++;
		} else if(hands.cards[i].face == "D") {
			D++;
		} else if(hands.cards[i].face == "C") {
			C++;
		}
	}

	ints := []int{S, H, D, C}

	sort.Ints(ints)
	return ints[3];
}

func IsFlsh(hands PokerHands) bool {
	return SameFace(hands) == 5;
}


// func IsStaightFlsh(cards [5]Card) {
// 	return true;
// }
