package main

import (
	"embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed inputs/*.txt
var inputFS embed.FS

func main() {
	input1, err := inputFS.ReadFile("inputs/input.txt")
	if err != nil {
		log.Panic(err)
	}
	pt1 := part1(string(input1))
	fmt.Println("The Fisrt part result is:", pt1)
	input2, err := inputFS.ReadFile("inputs/test.txt")
	if err != nil {
		log.Panic(err)
	}
	pt2 := part2(string(input2))
	fmt.Println("The Second part result is:", pt2)
}

func part1(input string) int {
	var result int
	in := strings.Split(input, "\n")

	for _, i := range in {
		var line []int

		for _, j := range strings.Split(i, "") {
			num, err := strconv.Atoi(j)
			if err == nil {
				line = append(line, num)
			}
		}
		result += line[0]*10 + line[len(line)-1]
	}

	return result
}

func part2(input string) int {
	var result int
	numbers := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}
	in := strings.Split(input, "\n")
	for _, i := range in {
		for k, j := range strings.Split(i, "") {
			switch j {
			case "o":
			case "t":
				if j+strings.Split(i, "")[k+1]+strings.Split(i, "")[k+2] == "two" {
					fmt.Println(j + strings.Split(i, "")[k+1] + strings.Split(i, "")[k+2])
				}
			case "f":
			case "s":
			case "e":
			case "n":
			}
		}
	}
	fmt.Println(numbers[""])
	return result
}
