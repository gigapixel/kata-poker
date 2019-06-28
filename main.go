package main

import (
	"fmt"
	"sort"
	"strconv"
	// "reflect"
)

// Card -
type Card struct {
	points string // 2 - A
	face   string // S,H,D,C
}

// S - หน้าโพธิ์ดำ
// H - หัวใจ
// D - ข้าวหลามตัด
// C - ดอกจิก
// 2 3 4 - 10 J M K A

// PokerHands -
type PokerHands struct {
	cards [5]Card
}

func main() {
	// // FIVE
	// cards := [5]Card{
	// 	{"A", "S"}, 
	// 	{"A", "H"}, 
	// 	{"A", "D"}, 
	// 	{"A", "C"}, 
	// 	{"J", "J"},
	// }

	// // ROYAL
	// cards := [5]Card{
	// 	{"A", "S"}, 
	// 	{"K", "S"}, 
	// 	{"Q", "S"}, 
	// 	{"10", "S"}, 
	// 	{"J", "S"},
	// }

	// HIGH CARD
	cards := [5]Card{
		// {"A", "H"}, 
		// {"2", "S"}, 
		// {"7", "D"}, 
		// {"10", "S"}, 
		// {"J", "C"},
		Card{"3", "H"},
		Card{"2", "S"},
		Card{"7", "C"},
		Card{"10", "D"},
		Card{"3", "S"},
	}

	// // FOUR
	// cards := [5]Card{
	// 	{"A", "S"}, 
	// 	{"A", "H"}, 
	// 	{"A", "D"}, 
	// 	{"A", "C"}, 
	// 	{"J", "J"},
	// }

	// // Straight flush*
	// cards := [5]Card{
	// 	Card{"J", "S"}, 
	// 	Card{"10", "S"}, 
	// 	Card{"9", "S"}, 
	// 	Card{"8", "S"}, 
	// 	Card{"7", "S"},
	// }

	// // Flush
	// cards := [5]Card{
	// 	Card{"J", "D"}, 
	// 	Card{"9", "D"}, 
	// 	Card{"8", "D"}, 
	// 	Card{"4", "D"}, 
	// 	Card{"3", "D"},
	// }

	// // Full  house
	// cards := [5]Card{
	// 	Card{"6", "S"}, 
	// 	Card{"6", "D"}, 
	// 	Card{"6", "H"}, 
	// 	Card{"11", "C"}, 
	// 	Card{"11", "K"},
	// }

	// // Three of a kind
	// cards := [5]Card{
	// 	Card{"6", "S"}, 
	// 	Card{"6", "D"}, 
	// 	Card{"6", "H"}, 
	// 	Card{"9", "C"}, 
	// 	Card{"11", "K"},
	// }

	// // Two pair
	// cards := [5]Card{
	// 	Card{"6", "S"}, 
	// 	Card{"6", "D"}, 
	// 	Card{"9", "H"}, 
	// 	Card{"9", "C"}, 
	// 	Card{"11", "K"},
	// }

	// // One Pair
	// cards := [5]Card{
	// 	Card{"6", "S"}, 
	// 	Card{"6", "D"}, 
	// 	Card{"8", "H"}, 
	// 	Card{"9", "C"}, 
	// 	Card{"11", "K"},
	// }

	// // High card
	// cards := [5]Card{
	// 	Card{"5", "S"}, 
	// 	Card{"5", "D"}, 
	// 	Card{"5", "H"}, 
	// 	Card{"5", "C"}, 
	// 	Card{"2", "K"},
	// }

	hand := PokerHands{cards}
	var res = cal(hand);
	fmt.Println(res);
}

func cal(hand PokerHands) string {
	strs := ""
	var isStraight = isStraight(hand);
	switch os := true; os {
		case (isStraight == "Straight"):
			strs += "Straight";
		case (isStraight == "Royal"):
			strs += "RoyalStraight";
		default:
	}

	// flush
	var rankF = ranks(hand, "face");
	switch true {
		case (rankF == "5"):
			strs += "Flush";
		default:
	}


	switch true {
		case (isStraight == "Straight" || isStraight == "RoyalStraight" || (rankF == "5")):
			return strs;
		default:
	}

	var rank = ranks(hand, "points");
	switch true {
	case (rank == "4,1hasJA"):
		strs += "Five of a kind";
	case (rank == "4,1"):
		strs += "Four of a kind";
	case (rank == "3,2"):
		strs += "Full house";
	case (rank == "3,1,1"):
		strs += "Three of a kind";
	case (rank == "2,2,1"):
		strs += "Two pair";
	case (rank == "2,1,1,1"):
		strs += "One pair";
	case (rank == "1,1,1,1,1"):
		list := []Card{}
		for _, k := range hand.cards {
			list = append(list, k);
		}
		sort.SliceStable(list, func(i, j int) bool {
			return convertPoint(list[i].points) < convertPoint(list[j].points)
		})
		strs += "High card " + list[4].points + ":" + list[4].face;
	default:
	}

	return strs;
}

func convertPoint(points string) int64 {
	switch os := points; os {
	case "J":
		return 11;
	case "Q":
		return 12;
	case "K":
		return 13;
	case "A":
		return 14;
	default:
		point, err := strconv.ParseInt(points, 10, 64)
		switch true {
			case (err != nil) : 
			{}
		}
		return point;
	}
}

func isStraight(hand PokerHands) string {
	list := []Card{}
	for _, k := range hand.cards {
		list = append(list, k);
	}
	sort.SliceStable(list, func(i, j int) bool {
		return convertPoint(list[i].points) < convertPoint(list[j].points)
	})
	for idx, card := range list {
		switch true {
			case idx != 4:
				next := convertPoint(list[idx + 1].points);
				me := convertPoint(card.points) + 1;
				switch true {
					case (next != me):
						return "";
				}
				break;
		}
	}

	switch true {
		case (list[4].points == "A"):
			return "Royal";
		default:
			return "Straight";
	}
}

func ranks(hand PokerHands, rankType string) string {
	// GROUP
	m := make(map[string][]string);
	for _, card := range hand.cards {
		switch os := rankType; os {
		case "face":
			m[card.face] = append(m[card.face], card.face) 
			break;
		default:
			m[card.points] = append(m[card.points], card.points) 
			break;
		}
	}

	// COUNT SAME VALUE
	count := make([]int, 0)
	for index := range m {
		count = append(count, len(m[index]))
	}
	
	// sort.Ints(s)
	sort.Slice(count, func(i, j int) bool {
		return count[i] > count[j]
	});

	var str = "";
	for index, val := range count {
		str += strconv.Itoa(val);
		switch true {
			case index + 1 != len(count):
				str += ",";
		}
	}

	var jk = "";
	for _, mval := range m {
		switch true {
			case ((str == "4,1") && (mval[0] == "J" || mval[0] == "A")):
				jk = "hasJA"
		}
	}
	return str + jk;
}