package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("02.txt")
	if err != nil {
		fmt.Println("ERROR")
	}
	pt1 := part1(string(input))
	fmt.Println("The Fisrt part result is:", pt1)
	pt2 := part2(string(input))
	fmt.Println("The Second part result is:", pt2)
}

func part1(input string) int {
	horizontal, depth := 0, 0
	in := strings.Split(strings.Trim(input, "\n"), "\n")

	for i := 0; i < len(in); i++ {
		direction := strings.Split(in[i], " ")[0]
		movement, _ := strconv.Atoi(strings.Split(in[i], " ")[1])
		if direction == "forward" {
			horizontal = horizontal + movement
		} else if direction == "up" {
			depth = depth - movement
		} else if direction == "down" {
			depth = depth + movement
		}
	}
	return horizontal * depth
}

func part2(input string) int {
	horizontal, depth, aim := 0, 0, 0
	in := strings.Split(strings.Trim(input, "\n"), "\n")

	for i := 0; i < len(in); i++ {
		direction := strings.Split(in[i], " ")[0]
		movement, _ := strconv.Atoi(strings.Split(in[i], " ")[1])
		if direction == "forward" {
			horizontal = horizontal + movement
			depth = depth + aim*movement
		} else if direction == "up" {
			aim = aim - movement
		} else if direction == "down" {
			aim = aim + movement
		}
	}
	return horizontal * depth
}
