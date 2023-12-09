package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/edwingeng/deque/v2"
)

type Card struct {
	WinningNumbers []int
	CardNumbers    []int
}

func main() {
	part1()
	//part2()
}

func part1() {
	fileLines := readInputFileLines()

	cards := make([]Card, 0)
	for _, line := range fileLines {
		split := strings.Split(line, ": ")
		twoLists := strings.Split(split[1], " | ")

		winningNumbers := convertToInts(strings.Split(twoLists[0], " "))
		cardNumbers := convertToInts(strings.Split(twoLists[1], " "))

		cards = append(cards, Card{winningNumbers, cardNumbers})
	}

	points := 0
	for _, card := range cards {
		matchingNumbers := intersection(card.WinningNumbers, card.CardNumbers)

		points += int(math.Pow(2, float64(len(matchingNumbers)-1)))
	}

	fmt.Println("Result:", points)
}

// Part2 is not functioning
func part2() {
	fileLines := readInputFileLines()

	cards := make([]Card, 0)
	for _, line := range fileLines {
		split := strings.Split(line, ": ")
		twoLists := strings.Split(split[1], " | ")

		winningNumbers := convertToInts(strings.Split(twoLists[0], " "))
		cardNumbers := convertToInts(strings.Split(twoLists[1], " "))

		cards = append(cards, Card{winningNumbers, cardNumbers})
	}

	cardQueue := deque.NewDeque[int]()
	copies := 0
	for index, card := range cards {
		matchingNumbers := intersection(card.WinningNumbers, card.CardNumbers)

		if len(matchingNumbers) == 0 {
			continue
		}

		copies += len(matchingNumbers)
		for _, item := range newSlice(index+1, len(matchingNumbers), 1) {
			cardQueue.Enqueue(item)
		}

		for !cardQueue.IsEmpty() {
			cardId := cardQueue.Dequeue()
			copyCard := cards[cardId]
			matchingNumbers := intersection(copyCard.WinningNumbers, copyCard.CardNumbers)

			copies += len(matchingNumbers)

			if len(matchingNumbers) <= 0 {
				continue
			}

			for _, item := range newSlice(index+1, len(matchingNumbers), 1) {
				cardQueue.Enqueue(item)
			}
		}

	}

	copies += len(cards)

	fmt.Println("Result:", copies)
}

func newSlice(start, count, step int) []int {
	s := make([]int, count)
	for i := range s {
		s[i] = start
		start += step
	}
	return s
}

func intersection(list1 []int, list2 []int) []int {
	matches := make([]int, 0)
	for _, item1 := range list1 {
		for _, item2 := range list2 {
			if item1 == item2 {
				matches = append(matches, item1)
				break
			}
		}
	}
	return matches
}

func convertToInts(list []string) []int {
	newList := make([]int, 0)

	for _, item := range list {
		if item == "" || item == " " {
			continue
		}
		num, _ := strconv.Atoi(item)
		newList = append(newList, num)
	}
	return newList
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
