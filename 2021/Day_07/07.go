package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("07.txt")
	if err != nil {
		log.Panic(err)
	}
	pt1 := part1(string(input))
	fmt.Println("The Fisrt part result is: ", pt1)
	pt2 := part2(string(input))
	fmt.Println("The Second part result is: ", pt2)
}

func part1(input string) int {
	var res int
	var max int
	in := strArrToIntArr(strings.Split(input, ","))
	for i := 0; i < len(in); i++ {
		if in[i] > max {
			max = in[i]
		}
	}
	crabs := make(map[int]int, max)
	for i := 0; i <= max; i++ {
		crabs[i] = 0
	}
	for _, i := range in {
		crabs[i]++
	}
	for i := 0; i < max; i++ {
		distances := findDistance(max, i)
		var fuel int
		for j := 0; j < len(crabs); j++ {
			fuel += distances[j] * crabs[j]
		}
		if res == 0 {
			res = fuel
		} else {
			if res > fuel {
				res = fuel
			}
		}
	}

	return res
}

func part2(input string) int {
	var res int
	var max int
	in := strArrToIntArr(strings.Split(input, ","))
	for i := 0; i < len(in); i++ {
		if in[i] > max {
			max = in[i]
		}
	}
	crabs := make(map[int]int, max)
	for i := 0; i <= max; i++ {
		crabs[i] = 0
	}
	for _, i := range in {
		crabs[i]++
	}
	for i := 0; i < max; i++ {
		distances := findDistance(max, i)
		var fuel int
		for j := 0; j < len(crabs); j++ {
			var distance int
			for k := 1; k <= distances[j]; k++ {
				distance += k
			}
			fuel += distance * crabs[j]
		}
		if res == 0 {
			res = fuel
		} else {
			if res > fuel {
				res = fuel
			}
		}
	}

	return res
}

func findDistance(max int, depth int) map[int]int {
	distances := make(map[int]int, max)
	for i := 0; i <= max; i++ {
		distances[i] = int(math.Abs(float64(depth - i)))
	}
	return distances
}

func strArrToIntArr(input []string) []int {
	res := make([]int, len(input))
	for i, j := range input {
		res[i], _ = strconv.Atoi(j)
	}
	return res
}
