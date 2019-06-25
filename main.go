package main

import "fmt"
import "sort"
// import "strconv"

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
	c2 := Card{points: "4", face: "S"}
	c3 := Card{points: "4", face: "S"}
	c4 := Card{points: "4", face: "S"}
	c5 := Card{points: "5", face: "A"}

	pokerHands := PokerHands{};
	pokerHands.cards[0] = c1;
	pokerHands.cards[1] = c2;
	pokerHands.cards[2] = c3;
	pokerHands.cards[3] = c4;
	pokerHands.cards[4] = c5;
	fmt.Println(pokerHands)
	
	// r1 := SameFace(pokerHands);
	// fmt.Println("Same Face : ", r1);

	// if(IsFlsh(pokerHands)) {
	// 	fmt.Println("Flsh");
	// } else if(SameNumber(pokerHands) == 4) {
	// 	fmt.Println("Four of a kind");
	// }

	getOutput(pokerHands);

	// fmt.Println("aaaa" , SameNumber(pokerHands));
	// IsContinueNumber(pokerHands, 3);
	
}

func getOutput(hands PokerHands) {
	if(IsFlsh(hands)) {
		fmt.Println("Flsh");
	} else if(SameNumber(hands) == 4) {
		fmt.Println("Four of a kind");
	}
}

// func IsContinueNumber(hands PokerHands, n) {
// 	var score []int
// 	for i := 0; i < 5; i++ {
// 		if(hands.cards[i].points == "K") {
// 			score = append(score, 13);
// 		} else if(hands.cards[i].points == "Q") {
// 			score = append(score, 12);
// 		} else if(hands.cards[i].points == "J") {
// 			score = append(score, 11);
// 		} else {
// 			k, _ := strconv.Atoi(hands.cards[i].points)
//     	score = append(score, k);
// 		}
// 	}

// 	sort.Ints(score)

// 	if(score[0] )

// 	for i := 0; i < 5; i++ {
// 		fmt.Println("score : ", score[i]);
// 	}
// }

func SameNumber(hands PokerHands) int {
	var n2 = 0;
	var n3 = 0;
	var n4 = 0;
	var n5 = 0;
	var n6 = 0;
	var n7 = 0;
	var n8 = 0;
	var n9 = 0;
	var n10 = 0;
	var n11 = 0;
	var n12 = 0;
	var n13 = 0;
	var n14 = 0;
	
	for i := 0; i < 5; i++ {
		if(hands.cards[i].points == "2") {
			n2++;
		} else if(hands.cards[i].points == "3") {
			n3++;
		} else if(hands.cards[i].points == "4") {
			n4++;
		} else if(hands.cards[i].points == "5") {
			n5++;
		} else if(hands.cards[i].points == "6") {
			n6++;
		} else if(hands.cards[i].points == "7") {
			n7++;
		} else if(hands.cards[i].points == "8") {
			n8++;
		} else if(hands.cards[i].points == "9") {
			n9++;
		} else if(hands.cards[i].points == "10") {
			n10++;
		} else if(hands.cards[i].points == "11") {
			n11++;
		} else if(hands.cards[i].points == "12") {
			n12++;
		} else if(hands.cards[i].points == "13") {
			n13++;
		} else if(hands.cards[i].points == "14") {
			n14++;
		}
	}

	ints := []int{n2, n3, n4, n5, n6, n7, n8, n9, n10, n11, n12, n13, n14};

	sort.Ints(ints)
	return ints[12];
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
