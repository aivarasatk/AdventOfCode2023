package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Input struct {
	Hand string
	Bid  int
}

var fileLines = make([]string, 0)

// integer strength value is only used for > < = operations. value 100 means greater than 99
var cardStrengthPart1 = map[rune]int{'A': 100, 'K': 99, 'Q': 98, 'J': 97, 'T': 96,
	'9': 95, '8': 94, '7': 93, '6': 92, '5': 91, '4': 90, '3': 89, '2': 87}

var cardStrengthPart2 = map[rune]int{'A': 100, 'K': 99, 'Q': 98, 'T': 96,
	'9': 95, '8': 94, '7': 93, '6': 92, '5': 91, '4': 90, '3': 89, '2': 87, 'J': 86}

func main() {
	part1()
	//PArt2 is ALMOST working :)
	//part2()
}

func part1() {
	fileLines := readInputFileLines()

	input := make([]Input, 0)
	for _, line := range fileLines {
		split := strings.Split(line, " ")
		bid, _ := strconv.Atoi(split[1])
		input = append(input, Input{split[0], bid})
	}

	highCards := make([]Input, 0)
	onePairCards := make([]Input, 0)
	twoPairCards := make([]Input, 0)
	threeOfKindCards := make([]Input, 0)
	fullHourseCards := make([]Input, 0)
	fourOfKindCards := make([]Input, 0)
	fiveOfKindCards := make([]Input, 0)

	for _, hand := range input {
		handAnalysis := map[rune]int{'A': 0, 'K': 0, 'Q': 0, 'J': 0, 'T': 0,
			'9': 0, '8': 0, '7': 0, '6': 0, '5': 0, '4': 0, '3': 0, '2': 0,
		}

		for _, card := range hand.Hand {
			handAnalysis[card]++
		}

		if isFiveOfAKind(handAnalysis) {
			fiveOfKindCards = append(fiveOfKindCards, hand)
		} else if isFourOfAKind(handAnalysis) {
			fourOfKindCards = append(fourOfKindCards, hand)
		} else if isFullHouse(handAnalysis) {
			fullHourseCards = append(fullHourseCards, hand)
		} else if isThreeOfAKind(handAnalysis) {
			threeOfKindCards = append(threeOfKindCards, hand)
		} else if isTwoPair(handAnalysis) {
			twoPairCards = append(twoPairCards, hand)
		} else if isOnePair(handAnalysis) {
			onePairCards = append(onePairCards, hand)
		} else if isHighCard(handAnalysis) {
			highCards = append(highCards, hand)
		}
	}

	orderedHands := make([]Input, 0)
	orderedHands = append(orderedHands, ordered(highCards, cardStrengthPart1)...)
	orderedHands = append(orderedHands, ordered(onePairCards, cardStrengthPart1)...)
	orderedHands = append(orderedHands, ordered(twoPairCards, cardStrengthPart1)...)
	orderedHands = append(orderedHands, ordered(threeOfKindCards, cardStrengthPart1)...)
	orderedHands = append(orderedHands, ordered(fullHourseCards, cardStrengthPart1)...)
	orderedHands = append(orderedHands, ordered(fourOfKindCards, cardStrengthPart1)...)
	orderedHands = append(orderedHands, ordered(fiveOfKindCards, cardStrengthPart1)...)

	var sum int64 = 0
	for i, hand := range orderedHands {
		sum += int64(hand.Bid * (i + 1))
	}

	fmt.Println("Result:", sum)
}

func part2() {
	fileLines := readInputFileLines()

	input := make([]Input, 0)
	for _, line := range fileLines {
		split := strings.Split(line, " ")
		bid, _ := strconv.Atoi(split[1])
		input = append(input, Input{split[0], bid})
	}

	highCards := make([]Input, 0)
	onePairCards := make([]Input, 0)
	twoPairCards := make([]Input, 0)
	threeOfKindCards := make([]Input, 0)
	fullHourseCards := make([]Input, 0)
	fourOfKindCards := make([]Input, 0)
	fiveOfKindCards := make([]Input, 0)

	for _, hand := range input {
		handAnalysis := map[rune]int{'A': 0, 'K': 0, 'Q': 0, 'J': 0, 'T': 0,
			'9': 0, '8': 0, '7': 0, '6': 0, '5': 0, '4': 0, '3': 0, '2': 0,
		}

		for _, card := range hand.Hand {
			handAnalysis[card]++
		}

		if isFiveOfAKindPart2(handAnalysis) {
			fiveOfKindCards = append(fiveOfKindCards, hand)
		} else if isFourOfAKindPart2(handAnalysis) {
			fourOfKindCards = append(fourOfKindCards, hand)
		} else if isFullHousePart2(handAnalysis) {
			fullHourseCards = append(fullHourseCards, hand)
		} else if isThreeOfAKindPart2(handAnalysis) {
			threeOfKindCards = append(threeOfKindCards, hand)
		} else if isTwoPairPart2(handAnalysis) {
			twoPairCards = append(twoPairCards, hand)
		} else if isOnePairPart2(handAnalysis) {
			onePairCards = append(onePairCards, hand)
		} else if isHighCard(handAnalysis) {
			highCards = append(highCards, hand)
		} else {
			fmt.Print()
		}
	}

	orderedHands := make([]Input, 0)
	orderedHands = append(orderedHands, ordered(highCards, cardStrengthPart2)...)
	orderedHands = append(orderedHands, ordered(onePairCards, cardStrengthPart2)...)
	orderedHands = append(orderedHands, ordered(twoPairCards, cardStrengthPart2)...)
	orderedHands = append(orderedHands, ordered(threeOfKindCards, cardStrengthPart2)...)
	orderedHands = append(orderedHands, ordered(fullHourseCards, cardStrengthPart2)...)
	orderedHands = append(orderedHands, ordered(fourOfKindCards, cardStrengthPart2)...)
	orderedHands = append(orderedHands, ordered(fiveOfKindCards, cardStrengthPart2)...)

	var sum int64 = 0
	for i, hand := range orderedHands {
		sum += int64(hand.Bid * (i + 1))
	}

	fmt.Println("Result:", sum)
}

func ordered(hand []Input, cardStrength map[rune]int) []Input {
	slices.SortFunc(hand,
		func(a, b Input) int {
			return compareHands(a.Hand, b.Hand, cardStrength)
		})

	return hand
}

func compareHands(a, b string, cardStrength map[rune]int) int {
	for index, char := range a {
		if cardStrength[char] < cardStrength[rune(b[index])] {
			return -1
		} else if cardStrength[char] > cardStrength[rune(b[index])] {
			return 1
		}
	}
	return 0
}

func isFiveOfAKind(handAnalysis map[rune]int) bool {
	for _, count := range handAnalysis {
		if count == 5 {
			return true
		}
	}
	return false
}

func isFourOfAKind(handAnalysis map[rune]int) bool {
	for _, count := range handAnalysis {
		if count == 4 {
			return true
		}
	}
	return false
}

func isFullHouse(handAnalysis map[rune]int) bool {
	hitsType1 := 0
	hitsType2 := 0
	for _, count := range handAnalysis {
		if count == 3 {
			hitsType1++
		}
		if count == 2 {
			hitsType2++
		}
	}
	return hitsType1 == 1 && hitsType2 == 1
}

func isThreeOfAKind(handAnalysis map[rune]int) bool {
	hitsType1 := 0
	hitsType2 := 0
	for _, count := range handAnalysis {
		if count == 3 {
			hitsType1++
		}
		if count == 1 {
			hitsType2++
		}
	}
	return hitsType1 == 1 && hitsType2 == 2
}

func isTwoPair(handAnalysis map[rune]int) bool {
	hitsType1 := 0
	hitsType2 := 0
	for _, count := range handAnalysis {
		if count == 2 {
			hitsType1++
		}
		if count == 1 {
			hitsType2++
		}
	}
	return hitsType1 == 2 && hitsType2 == 1
}

func isOnePair(handAnalysis map[rune]int) bool {
	hitsType1 := 0
	hitsType2 := 0
	for _, count := range handAnalysis {
		if count == 2 {
			hitsType1++
		}
		if count == 1 {
			hitsType2++
		}
	}
	return hitsType1 == 1 && hitsType2 == 3
}

func isHighCard(handAnalysis map[rune]int) bool {
	hits := 0
	for _, count := range handAnalysis {
		if count == 1 {
			hits++
		}
	}
	return hits == 5
}

func isFiveOfAKindPart2(handAnalysis map[rune]int) bool {
	for _, count := range handAnalysis {
		if count == 5 || (handAnalysis['J']+count == 5) {
			return true
		}
	}
	return false
}

func isFourOfAKindPart2(handAnalysis map[rune]int) bool {
	for key, count := range handAnalysis {
		if count == 4 || (handAnalysis['J']+count == 4 && key != 'J') {
			return true
		}
	}
	return false
}

func isFullHousePart2(handAnalysis map[rune]int) bool {
	hitsType1 := 0
	hitsType2 := 0
	jokerCounted := false
	for key, count := range handAnalysis {
		if count == 3 {
			hitsType1++
		} else if handAnalysis['J']+count == 3 && key != 'J' && !jokerCounted {
			hitsType1++
			jokerCounted = true
		} else if count == 2 {
			hitsType2++
		} else if handAnalysis['J']+count == 2 && key != 'J' && !jokerCounted {
			hitsType2++
			jokerCounted = true
		}
	}
	return hitsType1 == 1 && hitsType2 == 1
}

func isThreeOfAKindPart2(handAnalysis map[rune]int) bool {
	hitsType1 := 0
	hitsType2 := 0
	jokerCounted := false
	for key, count := range handAnalysis {
		if count == 3 {
			hitsType1++
		} else if handAnalysis['J']+count == 3 && !jokerCounted {
			hitsType1++
			jokerCounted = true
		} else if count == 1 && key != 'J' {
			hitsType2++
		}
	}
	return hitsType1 == 1 && hitsType2 == 2
}

func isTwoPairPart2(handAnalysis map[rune]int) bool {
	hitsType1 := 0
	hitsType2 := 0
	jokerCounted := false
	for key, count := range handAnalysis {
		if count == 2 {
			hitsType1++
		} else if handAnalysis['J']+count == 2 && key != 'J' && !jokerCounted {
			hitsType1++
			jokerCounted = true
		} else if count == 1 && key != 'J' {
			hitsType2++
		}
	}
	return hitsType1 == 2 && hitsType2 == 1
}

func isOnePairPart2(handAnalysis map[rune]int) bool {
	hitsType1 := 0
	hitsType2 := 0
	jokerCounted := false
	for key, count := range handAnalysis {
		if count == 2 {
			hitsType1++
		} else if handAnalysis['J']+count == 2 && key != 'J' && !jokerCounted {
			hitsType1++
			jokerCounted = true
		} else if count == 1 && key != 'J' {
			hitsType2++
		}
	}
	return hitsType1 == 1 && hitsType2 == 3
}

func readInputFileLines() []string {
	lines := make([]string, 0)

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
