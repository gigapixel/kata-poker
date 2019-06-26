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

	// High card
	cards := [5]Card{
		Card{"3", "S"}, 
		Card{"6", "D"}, 
		Card{"8", "H"}, 
		Card{"9", "C"}, 
		Card{"11", "K"},
	}

	hand := PokerHands{cards}
	var res = cal(hand);
	fmt.Println(res);
}

func cal(hand PokerHands) string {
	strs := "-"
	var isStraight = isStraight(hand);
	if(isStraight){
		strs += "Straight";
	}

	// flush
	var isFlush = isFlush(ranks(hand, "face"));
	if(isFlush){
		strs += "Flush";
	}

	if(isStraight || isFlush) {
		return strs;
	}

	var rank = ranks(hand, "points");
	if(rank == "4,1"){
		strs += "Five of a kind";
	}

	if(rank == "3,2"){
		strs += "Full house";
	}

	if(rank == "3,1,1"){
		strs += "Three of a kind";
	}

	if(rank == "2,2,1"){
		strs += "Two pair";
	}

	if(rank == "2,1,1,1"){
		strs += "One pair";
	}

	if(rank == "1,1,1,1,1"){
		strs += "High card";
	}

	return strs;
}

func convertPoint(points string) int64 {
	switch os := points; os {
	case "J":
		return 11;
	case "M":
		return 12;
	case "K":
		return 13;
	case "A":
		return 14;
	default:
		point, err := strconv.ParseInt(points, 10, 64)
		if err != nil {}	
		return point;
	}
}

func isSameFace(hand PokerHands, expect int) bool {
	m := make(map[string]string);
	for _, card := range hand.cards {
			m[card.face] = card.face
	}

	if(len(m) == expect) {
		return true;
	} else {
		return false;
	}
}

func isStraight(hand PokerHands) bool {
	list := []Card{}
	for _, k := range hand.cards {
		list = append(list, k);
	}
	sort.SliceStable(list, func(i, j int) bool {
		return convertPoint(list[i].points) < convertPoint(list[j].points)
	})
	for idx, card := range list {
		if(idx != 4){
			next := convertPoint(list[idx + 1].points);
			me := convertPoint(card.points) + 1;
			if(next != me){
				return false;
			}
		}
	}
	return true;
}

func isFlush(rank string) bool{
	if(rank == "5"){
		return true
	}else {
		return false;
	}
}

func ranks(hand PokerHands, rankType string) string {
	// GROUP
	m := make(map[string][]string);
	for _, card := range hand.cards {
		if(rankType == "face"){
			m[card.face] = append(m[card.face], card.face) 
		} else {
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
		if(index + 1 != len(count)){
			str += ",";
		}
	}
	// fmt.Println(str);
	return str;
}

