package main

import "fmt"
// import "strconv"
import "github.com/bradfitz/slice"

var m_point = map[string]int{
	"2": 2, 
	"3": 3, 
	"4": 4, 
	"5": 5, 
	"6": 6, 
	"7": 7, 
	"8": 8, 
	"9": 9, 
	"10": 10, 
	"J": 11, 
	"Q": 12, 
	"K": 13, 
	"A": 14,
}

var m_mapBool = map[bool]int {
	true: 1,
	false: 0,
}

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
	return m_point[c.points];
}

func main() {
	
	card1 := Card{"10", "C"}
	card2 := Card{"10", "C"}
	card3 := Card{"9", "C"}
	card4 := Card{"8", "C"}
	card5 := Card{"7", "C"}
	cards := []Card{card1, card2, card3, card4, card5}

	pokerHands := PokerHands{cards}

	fmt.Println("Cards: ", cards);
	
	result := printHand(pokerHands)
	fmt.Print(result)
	
}

func printHand(hand PokerHands) string {
	slice.Sort(hand.cards[:], func(i, j int) bool {
		return hand.cards[i].getPoint() < hand.cards[j].getPoint()
	})

	result := ""

	var customDistinct = getDistinct(hand.cards)
	var cntDistinct = len(customDistinct)
	var lenFirstElement = customDistinct[0];

	var isStraight = checkIsStraight(hand.cards); // int
	var isFlush = checkIsSameAllFaces(hand.cards); // int

	switch {
		case isStraight && isFlush:
			result = "Straight flush"
		case isStraight:
			result = "Straight"
		case isFlush:
			result = "Flush"
		case cntDistinct == 2 && lenFirstElement == 4:
			result = "Four of a kind"
		case cntDistinct == 2:
			result = "Full house"
		case cntDistinct == 3 && lenFirstElement == 3:
			result = "Three of a kind"
		case cntDistinct == 3:
			result = "Two pair"
		case cntDistinct == 4:
			result = "One pair"
		default:
			result = "High card"
	}

	return result
}

func getDistinct(cards []Card) []int {
	var s []int
	var cnt = 1
	var strPoint = cards[0].points

	for i := 1; i < len(cards); i++ {
		switch {
			case strPoint != cards[i].points:
				strPoint = cards[i].points
				s = append(s, cnt)

				cnt = 1				
			default:
				cnt++				
		}		
	}

	s = append(s, cnt)

	sortInt(s)  // sort desc
	return s
}

// Sort by Desc
func sortInt(arr []int) {
	slice.Sort(arr[:], func(i, j int) bool {
		return arr[i] > arr[j]
	})
}

func checkIsSameAllFaces(cards []Card) bool {
	for i := 1; i < len(cards); i++ {
		switch true {
			case cards[0].face != cards[i].face:
				return false;
		}		
	}

	return true
}

func checkIsStraight(cards []Card) bool {
	for i := 1; i < len(cards); i++ {
		switch true {
			case cards[i-1].getPoint() != (cards[i].getPoint() - 1):
				return false
		}	
	}

	return true
}

