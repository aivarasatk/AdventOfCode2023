package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Color int

const (
	Red Color = iota
	Green
	Blue
)

type Cube struct {
	Count int
	Color Color
}

type Set struct {
	Cubes []Cube
}

type Game struct {
	Id   int
	Sets []Set
}

var fileLines = make([]string, 0)

func main() {
	part1()
	part2()
}

// Part 1 just wants us read data correctly
func part1() {
	fileLines = readInputFileLines()

	games := make([]Game, 0)
	for _, line := range fileLines {
		splits := strings.Split(line, ": ") // space included in all splits to simplify parsing

		gameId, _ := strconv.Atoi(splits[0][5:])

		sets := strings.Split(splits[1], "; ")

		gameSets := make([]Set, 0)
		for _, set := range sets {
			cubes := strings.Split(set, ", ")
			cubesInSet := make([]Cube, 0)

			for _, cube := range cubes {
				pieces := strings.Split(cube, " ")
				count, _ := strconv.Atoi(pieces[0])
				color := Red

				if pieces[1] == "green" {
					color = Green
				} else if pieces[1] == "blue" {
					color = Blue
				}

				cubesInSet = append(cubesInSet, Cube{count, color})
			}

			gameSets = append(gameSets, Set{cubesInSet})
		}

		games = append(games, Game{gameId, gameSets})
	}

	gameIdSum := 0
	for _, game := range games {
		gameIdSum += game.Id
		for _, set := range game.Sets {
			redCubeCount := 12
			greenCubeCount := 13
			blueCubeCount := 14
			for _, cube := range set.Cubes {
				switch cube.Color {
				case Red:
					redCubeCount -= cube.Count
				case Blue:
					blueCubeCount -= cube.Count
				default:
					greenCubeCount -= cube.Count
				}
			}

			if redCubeCount < 0 || greenCubeCount < 0 || blueCubeCount < 0 {
				gameIdSum -= game.Id
				break
			}
		}
	}
	fmt.Println("Result:", gameIdSum)
}

// Part 2 wants to count how many cubes are needed to make each game possible
func part2() {
	fileLines = readInputFileLines()

	games := make([]Game, 0)
	for _, line := range fileLines {
		splits := strings.Split(line, ": ") // space included in all splits to simplify parsing

		gameId, _ := strconv.Atoi(splits[0][5:])

		sets := strings.Split(splits[1], "; ")

		gameSets := make([]Set, 0)
		for _, set := range sets {
			cubes := strings.Split(set, ", ")
			cubesInSet := make([]Cube, 0)

			for _, cube := range cubes {
				pieces := strings.Split(cube, " ")
				count, _ := strconv.Atoi(pieces[0])
				color := Red

				if pieces[1] == "green" {
					color = Green
				} else if pieces[1] == "blue" {
					color = Blue
				}

				cubesInSet = append(cubesInSet, Cube{count, color})
			}

			gameSets = append(gameSets, Set{cubesInSet})
		}

		games = append(games, Game{gameId, gameSets})
	}

	sum := 0
	for _, game := range games {
		minRedCubesNeeded := 1
		minGreenCubesNeeded := 1
		minBlueCubesNeeded := 1
		for _, set := range game.Sets {
			for _, cube := range set.Cubes {
				switch cube.Color {
				case Red:
					if minRedCubesNeeded < cube.Count {
						minRedCubesNeeded = cube.Count
					}
				case Blue:
					if minBlueCubesNeeded < cube.Count {
						minBlueCubesNeeded = cube.Count
					}
				default:
					if minGreenCubesNeeded < cube.Count {
						minGreenCubesNeeded = cube.Count
					}
				}
			}

		}
		sum += minRedCubesNeeded * minGreenCubesNeeded * minBlueCubesNeeded
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
