/* package main 

import "fmt" 

// Card - 
type Card struct { 
points string // 2 - A 
face string // S,H,D,C 
} 

// PokerHands - 
type PokerHands struct { 
cards [5]Card 
} 

func getOutput(hand PokerHands) string { 
for i := 0; i < 5; i++ { 
fmt.Println(hand.cards[i].points); 
} 
return "is" 
} 

func main() { 
fmt.Println("kata poker") 
cards := [5]Card{Card{"2", "H"}, Card{"2", "H"}, Card{"2", "H"}, Card{"2", "H"}, Card{"2", "H"}} 
hands := PokerHands{cards} 
output := getOutput(hands) 
// for i := 0; i < 5; i++ { 

// fmt.Println("hands" , hands.i) 
// } 

fmt.Println("output =>> " , output) 
} 

*/ 
package main 

import "fmt" 
import "sort" 
import "strconv" 

// Card - 
type Card struct { 
points string // 2 - A 
face string // S,H,D,C 
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
c4 := Card{points: "5", face: "S"} 
c5 := Card{points: "5", face: "A"} 

pokerHands := PokerHands{}; 
pokerHands.cards[0] = c1; 
pokerHands.cards[1] = c2; 
pokerHands.cards[2] = c3; 
pokerHands.cards[3] = c4; 
pokerHands.cards[4] = c5; 
fmt.Println(pokerHands) 

getOutput(pokerHands); 
} 

func getOutput(hands PokerHands) string{ 

var output = ""; 
switch true { 
case IsStraight(hands) && IsFlsh(hands) && IsUniqueNumber(hands): 
output = "Straight flush"; 
case IsFlsh(hands): 
output = "Flush"; 
case SameNumber(hands, 12) == 4: 
output = "Four of a kind"; 
case isFullhouse(hands): 
output = "Full house"; 
case IsStraight(hands): 
output = "Straight"; 
case SameNumber(hands, 12) == 3: 
output = "Three of a kind"; 
case SameNumber(hands, 12) == 2 && SameNumber(hands, 11) == 2: 
output = "Two pair"; 
case SameNumber(hands, 12) == 2: 
output = "One pair"; 
default: 
output = "High card"; 
} 
return output; 
} 

func isFullhouse(hands PokerHands) bool { 
return SameNumber(hands, 12) == 3 && SameNumber(hands, 11) == 2; 
} 


func SameNumber(hands PokerHands, index int) int { 
var n2 = 0; 
var n3 = 0; 
var n4 = 0; 
var n5 = 0; 
var n6 = 0; 
var n7 = 0; 
var n8 = 0; 
var n9 = 0; 
var n10 = 0; 
var nJ = 0; 
var nK = 0; 
var nQ = 0; 
var nA = 0; 

for i := 0; i < 5; i++ { 
switch hands.cards[i].points { 
case "2": 
n2++; 
case "3": 
n3++; 
case "4": 
n4++; 
case "5": 
n5++; 
case "6": 
n6++; 
case "7": 
n7++; 
case "8": 
n8++; 
case "9": 
n9++; 
case "10": 
n10++; 
case "J": 
nJ++; 
case "K": 
nK++; 
case "Q": 
nQ++; 
case "A": 
nA++; 
} 
} 

ints := []int{n2, n3, n4, n5, n6, n7, n8, n9, n10, nJ, nK, nQ, nA}; 

sort.Ints(ints) 
return ints[index]; 
} 

func SameFace(hands PokerHands) int { 
var S = 0; 
var H = 0; 
var D = 0; 
var C = 0; 

for i := 0; i < 5; i++ { 
switch hands.cards[i].face { 
case "S": 
S++; 
case "H": 
H++; 
case "D": 
D++; 
case "C": 
C++; 
} 
} 

ints := []int{S, H, D, C} 

sort.Ints(ints) 
return ints[3]; 
} 

func IsFiveCardContinueNumber(hands PokerHands) bool {
var score []int 
for i := 0; i < 5; i++ { 
switch hands.cards[i].points { 
case "K": 
score = append(score, 13); 
case "Q": 
score = append(score, 12); 
case "J": 
score = append(score, 11); 
default: 
k, _ := strconv.Atoi(hands.cards[i].points) 
score = append(score, k); 
} 
} 
sort.Ints(score) 
return score[0] + 4 == score[4]; 
} 


func IsFlsh(hands PokerHands) bool { 
return SameFace(hands) == 5; 
} 

func IsUniqueNumber(hands PokerHands) bool { 
return (SameNumber(hands, 0) <= 1) && (SameNumber(hands, 4) <= 1); 
} 

func IsStraight(hands PokerHands) bool { 
return (IsFiveCardContinueNumber(hands)); 
}