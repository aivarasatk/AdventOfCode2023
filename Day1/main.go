package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var fileLines = make([]string, 0)

func main() {
	part1()
	part2()
}

// Part 1 is implemented very naively with character individual checks for digit values
func part1() {
	fileLines = readInputFileLines()
	sum := 0

	for _, line := range fileLines {
		var firstNum, secondNum rune
		for i, char := range line {
			if firstNum == 0 && unicode.IsDigit(char) {
				firstNum = char
			}

			if firstNum != 0 && unicode.IsDigit(char) {
				secondNum = char
			}

			if i == len(line)-1 && firstNum == 0 {
				firstNum = char
				secondNum = firstNum
			} else if i == len(line)-1 && secondNum == 0 {
				secondNum = firstNum
			}
		}

		num, _ := strconv.Atoi(string(firstNum) + string(secondNum))
		sum += num
	}
	fmt.Println("Result:", sum)
}

// Part 1 uses known values as a map for search. finding first and last indices of key occurences
func part2() {
	fileLines = readInputFileLines()
	sum := 0

	digits := map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for _, line := range fileLines {

		var firstIndex = 99999999
		var firstValue = 0

		var lastIndex = -1
		var lastValue = 0

		for key, value := range digits {

			if index := strings.Index(line, key); index != -1 && index < firstIndex {
				firstIndex = index
				firstValue = value
			}

			if index := strings.LastIndex(line, key); index > lastIndex {
				lastIndex = index
				lastValue = value
			}
		}

		sum += firstValue*10 + lastValue
	}
	fmt.Println("Result:", sum)
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
