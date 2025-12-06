package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed inputs/01.txt
var inputs string

func main() {
	part1(inputs)
	part2(inputs)
}

func part1(input string) {
	lines := strings.Split(input, "\n")
	startingPos := 50
	passwd := 0
	for _, line := range lines {
		dir := string(line[0])
		mov, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		switch dir {
		case "L":
			startingPos -= mov
		case "R":
			startingPos += mov
		}

		if int(startingPos/100) != 0 && startingPos%100 == 0 || startingPos == 0 {
			passwd++
		}
	}

	fmt.Printf("Part1: %d\n", passwd)
}

func part2(input string) {
	lines := strings.Split(input, "\n")
	startingPos := 50
	passwd := 0

	for _, line := range lines {
		dir := string(line[0])
		mov, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		fmt.Printf("Starting pos: %d\n", startingPos)
		fmt.Printf("Mov: %d\n", mov)
		switch dir {
		case "L":
			if startingPos == 0 {
				startingPos = 100
			}
			startingPos -= mov
		case "R":
			if startingPos == 100 {
				startingPos = 0
			}
			startingPos += mov
		}

		fmt.Printf("After move: %d\n", startingPos)

		if startingPos == 0 || startingPos == 100 {
			passwd++
		}
		if startingPos < 0 {
			passwd++
			completedLoops := int(-startingPos / 100)
			remainder := -startingPos % 100
			passwd += completedLoops
			startingPos = 100 - remainder
		}

		if startingPos > 100 {
			completedLoops := int(startingPos / 100)
			remainder := startingPos % 100
			passwd += completedLoops
			startingPos = remainder
		}

		fmt.Printf("Final pos: %d\n", startingPos)
		fmt.Printf("Passwd: %d\n", passwd)
		fmt.Println("----")

	}
	fmt.Printf("Part2: %d\n", passwd)
}
