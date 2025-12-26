package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type ElementInList int

const (
	Ranges = iota
	Blank
	IngredientsId
)

//go:embed inputs/test.txt
var inputs string

func main() {
	part1(inputs)
	part2(inputs)
}

func part1(input string) {
	res := 0
	start, stop, id := returnIDRangesAndIngredientsId(input)
	if len(start) != len(stop) {
		panic("Start and Stop ranges are not the same length")
	}

	for _, i := range id {
		for j := 0; j < len(start); j++ {
			if i >= start[j] && i <= stop[j] {
				res++
				break
			}
		}
	}

	fmt.Printf("Part1: %d\n", res)
}

func part2(input string) {
	res := 0
	fmt.Printf("Part2: %d\n", res)
}

func returnIDRangesAndIngredientsId(input string) ([]int, []int, []int) {
	lines := strings.Split(input, "\n")
	listType := Ranges
	idRanges := make([]string, 0)
	ingredientsId := make([]string, 0)
	for _, line := range lines {
		if line == "" {
			listType = Blank
		}

		switch listType {
		case Ranges:
			idRanges = append(idRanges, line)
		case Blank:
			listType = IngredientsId
		case IngredientsId:
			ingredientsId = append(ingredientsId, line)
		}
	}

	startIdRangesInt := make([]int, 0)
	stopIdRangesInt := make([]int, 0)

	for _, idrange := range idRanges {
		ids := strings.Split(idrange, "-")
		start, err := strconv.Atoi(ids[0])
		if err != nil {
			panic(err)
		}
		stop, err := strconv.Atoi(ids[1])
		if err != nil {
			panic(err)
		}

		startIdRangesInt = append(startIdRangesInt, start)
		stopIdRangesInt = append(stopIdRangesInt, stop)
	}

	ingredientsIdInt := make([]int, 0)
	for _, ingredientId := range ingredientsId {
		id, err := strconv.Atoi(ingredientId)
		if err != nil {
			panic(err)
		}
		ingredientsIdInt = append(ingredientsIdInt, id)
	}

	return startIdRangesInt, stopIdRangesInt, ingredientsIdInt
}
