package main

import "fmt"

// Card -
type Card struct {
	points string // 2 - A
	face   string // S,H,D,C
}

// S - หน้าโพธิ์ดำ
// H - หัวใจ
// D - ข้าวหลามตัด
// C - ดอกจิก

// PokerHands -
type PokerHands struct {
	cards [5]Card
}

func main() {
	fmt.Println("kata poker");
	cards := [5]Card{
		Card{"A", "S"}, 
		Card{"A", "C"}, 
		Card{"A", "H"}, 
		Card{"A", "D"}, 
		Card{"", ""}
	}
	hand := PokerHands{cards}
	fmt.Println(hand);
}

func cal() {


	
}

