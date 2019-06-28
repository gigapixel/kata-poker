package main

import (
	"fmt"
	"sort"
	"strconv"
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

	hand := PokerHands{cards}
	var res = cal(hand);
	fmt.Println(res);
}

func cal(hand PokerHands) string {
	strs := ""
	var isStraight = isStraight(hand);
	switch os := true; os {
		case (isStraight != ""):
			strs += isStraight;
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
		case (isStraight != "" || (rankF == "5")):
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
		}
	}

	switch true {
		case (list[4].points == "A"):
			return "RoyalStraight";
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
		default:
			m[card.points] = append(m[card.points], card.points) 
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