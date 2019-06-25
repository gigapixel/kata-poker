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

	cards1 := [5]Card{Card{"9", "C"}, Card{"J", "C"}, Card{"10", "C"}, Card{"Q", "C"}, Card{"K", "C"}}
	numbers := []int{ 0, 0, 0, 0, 0 }
	faces := []string{ "", "", "", "", "" }
	for i := 0; i < 5; i++ {
		numbers[i] = convertStrToInt(cards1[i].points)
		faces[i] = cards1[i].face
	}
	sort.Ints(numbers)

	isStraight := checkStraight(numbers))
	fmt.Println(numbers)
}

func checkStraight(numbers []int) bool {
	for i, num := range numbers {
		if i == 0 {
			continue
		}
		if (numbers[i - 1] + 1) != num {
			return false
		}
	}

	return true
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
