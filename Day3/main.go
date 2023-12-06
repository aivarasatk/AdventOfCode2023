package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	Row    int
	Column int
}

func main() {
	part1()
	part2()
}

func part1() {
	lines := readInputFileLines()

	fileLines := make([][]rune, 0)

	for row, line := range lines {
		fileLines = append(fileLines, make([]rune, 0))
		for _, char := range line {
			fileLines[row] = append(fileLines[row], char)
		}
	}

	maxRowIndex := len(lines) - 1
	maxColumnIndex := len(lines[0]) - 1

	var sum int64 = 0
	for row, line := range fileLines {
		for column, character := range line {
			if !isSymbol(character) {
				continue
			}

			checkList := pointsToCheck(row, column, maxRowIndex, maxColumnIndex)

			for _, point := range checkList {
				if fileLines[point.Row][point.Column] < '0' || fileLines[point.Row][point.Column] > '9' {
					continue
				}

				leftMostDigitColumn := point.Column
				rightMostDigitColumn := point.Column

				for leftMostDigitColumn >= 0 && fileLines[point.Row][leftMostDigitColumn] >= '0' && fileLines[point.Row][leftMostDigitColumn] <= '9' {
					leftMostDigitColumn--
				}

				for rightMostDigitColumn <= maxColumnIndex && fileLines[point.Row][rightMostDigitColumn] >= '0' && fileLines[point.Row][rightMostDigitColumn] <= '9' {
					rightMostDigitColumn++
				}

				number, _ := strconv.ParseInt(string(fileLines[point.Row][leftMostDigitColumn+1:rightMostDigitColumn]), 10, 0)
				for i := leftMostDigitColumn + 1; i <= rightMostDigitColumn-1; i++ {
					fileLines[point.Row][i] = '.'
				}
				sum += number
			}

		}
	}
	fmt.Println("Result:", sum)
}

func part2() {
	lines := readInputFileLines()

	fileLines := make([][]rune, 0)

	for row, line := range lines {
		fileLines = append(fileLines, make([]rune, 0))
		for _, char := range line {
			fileLines[row] = append(fileLines[row], char)
		}
	}

	maxRowIndex := len(lines) - 1
	maxColumnIndex := len(lines[0]) - 1

	var sum int64 = 0
	for row, line := range fileLines {
		for column, character := range line {
			if !isPotentialGear(character) {
				continue
			}

			checkList := pointsToCheck(row, column, maxRowIndex, maxColumnIndex)

			gearAdjecentNumberCount := 0
			var ratio int64 = 1
			for _, point := range checkList {
				if fileLines[point.Row][point.Column] < '0' || fileLines[point.Row][point.Column] > '9' {
					continue
				}

				leftMostDigitColumn := point.Column
				rightMostDigitColumn := point.Column

				for leftMostDigitColumn >= 0 && fileLines[point.Row][leftMostDigitColumn] >= '0' && fileLines[point.Row][leftMostDigitColumn] <= '9' {
					leftMostDigitColumn--
				}

				for rightMostDigitColumn <= maxColumnIndex && fileLines[point.Row][rightMostDigitColumn] >= '0' && fileLines[point.Row][rightMostDigitColumn] <= '9' {
					rightMostDigitColumn++
				}

				number, _ := strconv.ParseInt(string(fileLines[point.Row][leftMostDigitColumn+1:rightMostDigitColumn]), 10, 0)
				for i := leftMostDigitColumn + 1; i <= rightMostDigitColumn-1; i++ {
					fileLines[point.Row][i] = '.'
				}
				gearAdjecentNumberCount++
				ratio *= number
			}

			if gearAdjecentNumberCount == 2 {
				sum += ratio
			}

		}
	}
	fmt.Println("Result:", sum)
}

func isSymbol(character rune) bool {
	return (character < '0' || character > '9') && character != '.'
}

func isPotentialGear(character rune) bool {
	return character == '*'
}

func pointsToCheck(row, column, maxRowIndex, maxColumnIndex int) []Point {
	checkList := make([]Point, 0)

	if row-1 >= 0 {
		checkList = append(checkList, Point{row - 1, column})
	}

	if row+1 <= maxRowIndex {
		checkList = append(checkList, Point{row + 1, column})
	}

	if column-1 >= 0 {
		checkList = append(checkList, Point{row, column - 1})
	}

	if column+1 <= maxColumnIndex {
		checkList = append(checkList, Point{row, column + 1})
	}

	//diagonal

	if row-1 >= 0 && column-1 >= 0 {
		checkList = append(checkList, Point{row - 1, column - 1})
	}

	if row-1 >= 0 && column+1 <= maxColumnIndex {
		checkList = append(checkList, Point{row - 1, column + 1})
	}

	if row+1 <= maxRowIndex && column-1 >= 0 {
		checkList = append(checkList, Point{row + 1, column - 1})
	}

	if row+1 <= maxRowIndex && column+1 <= maxColumnIndex {
		checkList = append(checkList, Point{row + 1, column + 1})
	}

	return checkList
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
