package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed inputs/04.txt
var inputs string

func main() {
	part1(inputs)
	part2(inputs)
}

func part1(input string) {
	var res int
	in := strings.Split(input, "\n")
	rolls := make([][]bool, 0)
	for _, line := range in {
		roll := make([]bool, 0)
		for l := range line {
			roll = append(roll, string(line[l]) == "@")
		}
		rolls = append(rolls, roll)
	}

	for r := range rolls {
		for c := range rolls[r] {
			adj := checkAdjacentsRolls(rolls, r, c)
			if adj < 4 && rolls[r][c] {
				res++
			}
		}
	}

	fmt.Printf("Part1: %d\n", res)
}

func part2(input string) {
	var res int
	fmt.Printf("Part2: %d\n", res)
}

func checkAdjacentsRolls(rolls [][]bool, row, col int) int {
	var adj int
	if !rolls[row][col] {
		return adj
	}

	cols := len(rolls[0])
	rows := len(rolls)
	rowsToCheck := make([]int, 0)
	for r := row - 1; r <= row+1; r++ {
		if r >= 0 && r < rows {
			rowsToCheck = append(rowsToCheck, r)
		}
	}

	colsToCheck := make([]int, 0)
	for c := col - 1; c <= col+1; c++ {
		if c >= 0 && c < cols {
			colsToCheck = append(colsToCheck, c)
		}
	}

	for _, r := range rowsToCheck {
		for _, c := range colsToCheck {
			if r == row && c == col {
				continue
			}

			if rolls[r][c] {
				adj++
			}
		}
	}
	return adj
}
