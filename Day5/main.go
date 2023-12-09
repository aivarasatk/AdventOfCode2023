package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var fileLines = make([]string, 0)

type Range struct {
	Destination int64
	Source      int64
	Length      int64
}

type Chunk struct {
	Start int64
	End   int64
}

func main() {
	part1()
	part2()
}

func part1() {
	fileLines = readInputFileLines()

	source := make([]int64, 0)
	for _, item := range strings.Split(strings.TrimLeft(fileLines[0], "seeds: "), " ") {
		num, _ := strconv.ParseInt(item, 10, 64)
		source = append(source, num)
	}

	i := 3
	maps := make([][]Range, 0)
	ranges := make([]Range, 0)
	for i < len(fileLines) {
		if fileLines[i] == "" {
			i += 2
			maps = append(maps, ranges)

			ranges = nil
			ranges = make([]Range, 0)
		}

		row := strings.Split(fileLines[i], " ")

		destination, _ := strconv.ParseInt(row[0], 10, 64)
		source, _ := strconv.ParseInt(row[1], 10, 64)
		length, _ := strconv.ParseInt(row[2], 10, 64)

		ranges = append(ranges, Range{destination, source, length})

		i++
	}
	maps = append(maps, ranges)

	var destination []int64
	for _, ranges := range maps {
		destination = nil
		destination = make([]int64, 0)
		sourceUsed := make([]int64, 0)
		for _, item := range ranges {
			for _, src := range source {
				if src >= item.Source && src <= item.Source+item.Length {
					sourceUsed = append(sourceUsed, src)
					destination = append(destination, item.Destination+src-item.Source)
				}
			}
		}

		destination = append(destination, excluding(source, sourceUsed)...)
		source = nil
		source = append(source, destination...)
	}

	fmt.Println("Result:", slices.Min(destination))
}

func part2() {

}

func excluding(source []int64, excludeList []int64) []int64 {
	matches := make([]int64, 0)
	for _, item1 := range source {
		found := false
		for _, item2 := range excludeList {
			if item1 == item2 {
				found = true
				break
			}
		}
		if !found {
			matches = append(matches, item1)
		}
	}
	return matches
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
