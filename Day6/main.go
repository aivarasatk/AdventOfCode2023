package main

import (
	"bufio"
	"fmt"
	"os"
)

var fileLines = make([]string, 0)

type Race struct {
	TimeGiven           int64
	MaxTraveledDistance int64
}

func main() {
	part1()
	part2()
}

func part1() {

	/*
		Time:      7  15   30
		Distance:  9  40  200
	*/
	// races := [3]Race{
	// 	{7, 9},
	// 	{15, 40},
	// 	{30, 200},
	// }

	/*
		Time:        56     97     78     75
		Distance:   546   1927   1131   1139
	*/
	races := [4]Race{
		{56, 546},
		{97, 1927},
		{78, 1131},
		{75, 1139},
	}

	result := 1
	for _, race := range races {
		reasoanableStart := race.MaxTraveledDistance/race.TimeGiven + 1 // +1 becausse my input has denominators always
		maximumumReasonableHoldTime := race.TimeGiven - reasoanableStart
		possibleWins := 0
		for i := reasoanableStart; i <= maximumumReasonableHoldTime; i++ {
			if i*(race.TimeGiven-i) > race.MaxTraveledDistance {
				possibleWins++
			}
		}
		result *= possibleWins
	}
	fmt.Println("Result:", result)
}

func part2() {
	/*
		Time:      71530
		Distance:  940200
	*/
	// race := Race{71530, 940200}
	/*
		Time:       56977875
		Distance:   546192711311139
	*/
	race := Race{56977875, 546192711311139}

	reasoanableStart := race.MaxTraveledDistance/race.TimeGiven + 1 // +1 becausse my input has denominators always
	maximumumReasonableHoldTime := race.TimeGiven - reasoanableStart
	possibleWins := 0
	for i := reasoanableStart; i <= maximumumReasonableHoldTime; i++ {
		if i*(race.TimeGiven-i) > race.MaxTraveledDistance {
			possibleWins++
		}
	}

	fmt.Println("Result:", possibleWins)
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
