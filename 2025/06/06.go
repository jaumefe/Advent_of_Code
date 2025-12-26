package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed inputs/06.txt
var inputs string

func main() {
	part1(inputs)
	part2(inputs)
}

func part1(input string) {
	res := 0
	operators, columns := parseInputs(input)
	for i, op := range operators {
		switch op {
		case "+":
			res += sum(columns[i])
		case "*":
			res += multiply(columns[i])
		}
	}
	fmt.Printf("Part1: %d\n", res)
}

func part2(input string) {
	res := 0

	fmt.Printf("Part2: %d\n", res)
}

func parseInputs(input string) (operators []string, columns [][]int) {
	lines := strings.Split(input, "\n")

	// Get operator's row
	operatorsRaw := lines[len(lines)-1]
	operators = make([]string, 0)
	for _, op := range operatorsRaw {
		if string(op) != " " {
			operators = append(operators, string(op))
		}
	}

	// Get numbers arranged per columns
	numbersRaw := lines[:len(lines)-1]
	numbersStr := make([][]string, 0)
	numLines := make([]string, 0)
	numTmp := ""
	for _, numLine := range numbersRaw {
		for i, numChar := range numLine {
			if string(numChar) != " " {
				numTmp += string(numChar)
			} else {
				if numTmp != "" {
					numLines = append(numLines, numTmp)
					numTmp = ""
				}
			}

			if i == len(numLine)-1 {
				if numTmp != "" {
					numLines = append(numLines, numTmp)
					numTmp = ""
				}
			}
		}
		numbersStr = append(numbersStr, numLines)
		numLines = make([]string, 0)
	}

	columns = make([][]int, 0)
	colLen := len(numbersStr[0])
	rowLen := len(numbersStr)

	for c := 0; c < colLen; c++ {
		col := make([]int, 0)
		for r := 0; r < rowLen; r++ {
			num, err := strconv.Atoi(numbersStr[r][c])
			if err != nil {
				panic(err)
			}
			col = append(col, num)
		}
		columns = append(columns, col)
	}

	return
}

func sum(column []int) int {
	res := 0
	for _, num := range column {
		res += num
	}
	return res
}

func multiply(column []int) int {
	res := 1
	for _, num := range column {
		res *= num
	}
	return res
}
