package main

import "fmt"
import "strconv"
import "sort"

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

	// validCardNumbers = ["2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"];
	// cardValues = [2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14];
	// validCardSuits = ["H", "D", "C", "S"];
	// var J=11, Q=12, K=13, A=14, C=1, D=2, H=4, S=8;

	cards1 := [5]Card{Card{"A", "C"}, Card{"7", "C"}, Card{"K", "C"}, Card{"J", "C"}, Card{"5", "C"}}
	numbers := []int{ 0, 0, 0, 0, 0}
	for i := 0; i < 5; i++ {
		numbers[i] = convertStrToInt(cards1[i].points)
	}
	sort.Ints(numbers)

	fmt.Println(numbers)
}

func checkStraight(numbers []int) bool {
	for ()
}

func convertStrToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err == nil {
		return i
	} else {
		switch str {
		case "J":
			return 11
		case "Q":
			return 12
		case "A":
			return 14
		case "K":
			return 13
		}
		return 0
	}
}
