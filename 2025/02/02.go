package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed inputs/02.txt
var inputs string

func main() {
	part1(inputs)
	part2(inputs)
}

func part1(input string) {
	ranges := strings.Split(input, ",")
	res := 0

	for _, r := range ranges {
		rangeParts := strings.Split(r, "-")
		firstStr, lastStr := rangeParts[0], rangeParts[1]

		if len(firstStr)%2 != 0 && len(lastStr)%2 != 0 && len(firstStr) == len(lastStr) {
			continue
		}

		first, err := strconv.Atoi(firstStr)
		if err != nil {
			panic(err)
		}

		last, err := strconv.Atoi(lastStr)
		if err != nil {
			panic(err)
		}

		for i := first; i <= last; i++ {
			iStr := strconv.Itoa(i)
			length := len(iStr)
			divisor := math.Pow(10, float64(length/2))
			if int(i/int(divisor)) == i%int(divisor) {
				res += i
			}

		}
	}
	fmt.Printf("Part1: %d\n", res)
}

func part2(input string) {

}
