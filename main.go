package main

import "fmt"
import "strconv"
import "github.com/bradfitz/slice"

// Card -
type Card struct {
	points string // 2 - A
	face   string // S,H,D,C
}

// PokerHands -
type PokerHands struct {
	cards []Card
}

func (c Card) getPoint() int {
	if c.points == "A" {
		return 14
	} else if c.points == "K" {
		return 13
	} else if c.points == "Q" {
		return 12
	} else if c.points == "J" {
		return 11
	}

	s, _ := strconv.ParseInt(c.points, 10, 32)
	return int(s)
}

func main() {
	// fmt.Println("kata poker")
	card1 := Card{"10", "C"}
	card2 := Card{"10", "C"}
	card3 := Card{"9", "C"}
	card4 := Card{"8", "C"}
	card5 := Card{"7", "C"}
	// hand := PokerHands(cards)
	// hand.print(card1)
	cards := []Card{card1, card2, card3, card4, card5}	

	pokerHands := PokerHands{cards}

	result := printHand(pokerHands)
	fmt.Print(result)
	//return result
}

func printHand(hand PokerHands) string {
	slice.Sort(hand.cards[:], func(i, j int) bool {
		return hand.cards[i].getPoint() < hand.cards[j].getPoint()
	})

	result := ""	

	// var cntDistinct = countDistinctPoint(hand.cards);

	var customDistinct = getDistinct(hand.cards); // 
	var cntDistinct = len(customDistinct);
	// fmt.Println("Count:", cntDistinct, " Data: ", customDistinct)

	/* ยังใช้ cntDistinct ตรงๆ ไม่ได้
		เพราะ Three of kind กับ Two Pair นับได้ 3 ใบเหมือนกัน
	*/
	if (StraightFlush(hand.cards)) {
		result = "Straight flush"
	} else if (cntDistinct == 2 && customDistinct[0] == 4 /* || FourOfAKind(hand.cards) */ ) {
		result = "Four of a kind"
	} else if (cntDistinct == 2 && customDistinct[0] == 3 /* FullHouse(hand.cards) */) {
		result = "Full house"
	} else if (Flush(hand.cards)) {
		result = "Flush"
	} else if (Straight(hand.cards)) {
		result = "Straight"
	}	else if (cntDistinct == 3 && customDistinct[0] == 3 /* ThreeOfAKind(hand.cards) */ ) {
		result = "Three of a kind"
	} else if (cntDistinct == 3 /* || TwoPair(hand.cards) */ ) {
		result = "Two pair"
	} else if (cntDistinct == 4 /* || OnePair(hand.cards) */ ) {
		result = "One pair"
	} else {
		result = "High card"
	}

	return result
}

// func countDistinctPoint(cards []Card) int {
// 	var cnt = 1;

// 	var strPoint = cards[0].points;
// 	for i := 1; i < len(cards); i++ {
// 		if strPoint != cards[i].points {
// 			strPoint = cards[i].points;
// 			cnt++
// 		}
// 	}
	
// 	return cnt;
// }

func getDistinct(cards []Card) []int {
	var s []int

	var cnt = 1;

	var strPoint = cards[0].points;
	for i := 1; i < len(cards); i++ {
		if strPoint != cards[i].points {
			strPoint = cards[i].points;
			s = append(s, cnt)

			cnt = 1
		} else {
			cnt++
		}
	}

	s = append(s, cnt)


	sortInt(s, false);	
	return s;
}

func sortInt(arr []int, sortAsc bool) /* []int */ {
	slice.Sort(arr[:], func(i, j int) bool {		
		if (sortAsc) {
			return arr[i] < arr[j]
		} else {
			return arr[i] > arr[j]
		}
	});

	// return arr;
}

func sameAllFaces(cards []Card) bool {
	for i := 1; i < len(cards); i++ {
		if cards[0].face != cards[i].face {
			return false
		}
	}

	return true
}


func StraightFlush(cards []Card) bool {
	pt0 := cards[0].getPoint()
	pt1 := cards[1].getPoint()
	pt2 := cards[2].getPoint()
	pt3 := cards[3].getPoint()
	pt4 := cards[4].getPoint()

	isCorrectPoint := pt0+1 == pt1 && pt0+2 == pt2 && pt0+3 == pt3 && pt0+4 == pt4
	return isCorrectPoint && sameAllFaces(cards)
}

func FourOfAKind(cards []Card) bool {
	return false;
}

func FullHouse(cards []Card) bool {
	return false;
}

func Straight(cards []Card) bool {
	pt0 := cards[0].getPoint()
	pt1 := cards[1].getPoint()
	pt2 := cards[2].getPoint()
	pt3 := cards[3].getPoint()
	pt4 := cards[4].getPoint()

	isCorrectPoint := pt0+1 == pt1 && pt0+2 == pt2 && pt0+3 == pt3 && pt0+4 == pt4
	return isCorrectPoint
}

func Flush(cards []Card) bool {
	return sameAllFaces(cards);
}

func ThreeOfAKind(cards []Card) bool {
	return false;
}

func TwoPair(cards []Card) bool {
	return false;
}

func OnePair(cards []Card) bool {
	return false;
}

// func FourOfKind(cards [5]Card) bool {
// 	var cntDup := 0;
// 	for i := 0; i < len(cards); i++ {

// 	}

// 	return false
// }
