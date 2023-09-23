package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("01.txt")
	if err != nil {
		fmt.Println("ERROR")
	}
	pt1 := part1(string(input))
	fmt.Println("The Fisrt part result is:", pt1)
	pt2 := part2(string(input))
	fmt.Println("The Second part result is:", pt2)
}

func part1(input string) int {
	i := 0
	depth := strings.Split(strings.Trim(input, "\n"), "\n")
	prev_depth, _ := strconv.Atoi(depth[0])
	for _, ndepth := range depth {
		deep, _ := strconv.Atoi(ndepth)
		if deep > prev_depth {
			i = i + 1
		}
		prev_depth = deep
	}
	return i
}

func part2(input string) int {
	result := 0
	prev_depth_sum := 0
	depth := strings.Split(strings.Trim(input, "\n"), "\n")
	for i := 0; i < len(depth)-3; i++ {
		deep1, _ := strconv.Atoi(depth[i])
		deep2, _ := strconv.Atoi(depth[i+1])
		deep3, _ := strconv.Atoi(depth[i+2])
		deep := deep1 + deep2 + deep3
		if deep > prev_depth_sum {
			result = result + 1
		}
		prev_depth_sum = deep
	}
	return result
}
