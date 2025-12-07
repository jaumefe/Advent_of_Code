package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed inputs/test.txt
var inputs string

func main() {
	part1(inputs)
	part2(inputs)
}

func part1(input string) {
	res := 0
	banks := strings.SplitSeq(input, "\n")

	for bank := range banks {
		b := strings.Split(bank, "")
		max1, pos := findMaxAndPos1(b)
		b = b[pos+1:]
		max2 := findMax(b)
		res += max1*10 + max2
	}
	fmt.Printf("Part1: %d\n", res)
}

func part2(input string) {
	res := 0
	banks := strings.SplitSeq(input, "\n")
	for bank := range banks {
		b := strings.Split(bank, "")
		pos := 0
		max := make([]int, 12)

		for i := range 11 {
			max[i], pos = findMaxAndPos2(b[pos:], 11-i)
			b = b[pos+1:]
			fmt.Println("Max:", max[i], "Pos:", pos)
		}

		max[11] = findMax(b)
		fmt.Println(max)
		for i, m := range max {
			res += m * int(math.Pow(10, float64(11-i)))
		}
	}
	fmt.Printf("Part2: %d\n", res)

}

func findMaxAndPos1(bank []string) (int, int) {
	max := 0
	pos := 0

	for i, b := range bank {
		val, err := strconv.Atoi(b)
		if err != nil {
			panic(err)
		}

		if val > max {
			if i != len(bank)-1 {
				max = val
				pos = i
			}
		}
	}
	return max, pos
}

func findMaxAndPos2(bank []string, stopLen int) (int, int) {
	max := 0
	pos := 0

	for i, b := range bank {
		val, err := strconv.Atoi(b)
		if err != nil {
			panic(err)
		}

		if i == len(bank)-stopLen {
			max = val
			pos = i
			return max, pos
		}

		if val > max {
			if i != len(bank)-1 {
				max = val
				pos = i
			}
		}
	}
	return max, pos
}

func findMax(bank []string) int {
	max := 0

	for _, b := range bank {
		val, err := strconv.Atoi(b)
		if err != nil {
			panic(err)
		}

		if val > max {
			max = val
		}
	}
	return max
}
