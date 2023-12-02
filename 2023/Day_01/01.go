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
	input2, err := inputFS.ReadFile("inputs/input.txt")
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
	in := strings.Split(input, "\n")
	for _, i := range in {
		var line []int
		for k, j := range strings.Split(i, "") {
			num, err := strconv.Atoi(j)
			if err == nil {
				line = append(line, num)
			} else {
				switch j {
				case "o":
					if k+2 < len(strings.Split(i, "")) {
						if j+strings.Split(i, "")[k+1]+strings.Split(i, "")[k+2] == "one" {
							line = append(line, 1)
						}
					}
				case "t":
					if k+2 < len(strings.Split(i, "")) {
						if j+strings.Split(i, "")[k+1]+strings.Split(i, "")[k+2] == "two" {
							line = append(line, 2)
						}
					}
					if k+4 < len(strings.Split(i, "")) {
						if j+strings.Split(i, "")[k+1]+strings.Split(i, "")[k+2]+strings.Split(i, "")[k+3]+strings.Split(i, "")[k+4] == "three" {
							line = append(line, 3)
						}
					}
				case "f":
					if k+3 < len(strings.Split(i, "")) {
						if j+strings.Split(i, "")[k+1]+strings.Split(i, "")[k+2]+strings.Split(i, "")[k+3] == "four" {
							line = append(line, 4)
						}
						if j+strings.Split(i, "")[k+1]+strings.Split(i, "")[k+2]+strings.Split(i, "")[k+3] == "five" {
							line = append(line, 5)
						}
					}
				case "s":
					if k+2 < len(strings.Split(i, "")) {
						if j+strings.Split(i, "")[k+1]+strings.Split(i, "")[k+2] == "six" {
							line = append(line, 6)
						}
					}
					if k+4 < len(strings.Split(i, "")) {
						if j+strings.Split(i, "")[k+1]+strings.Split(i, "")[k+2]+strings.Split(i, "")[k+3]+strings.Split(i, "")[k+4] == "seven" {
							line = append(line, 7)
						}
					}
				case "e":
					if k+4 < len(strings.Split(i, "")) {
						if j+strings.Split(i, "")[k+1]+strings.Split(i, "")[k+2]+strings.Split(i, "")[k+3]+strings.Split(i, "")[k+4] == "eight" {
							line = append(line, 8)
						}
					}
				case "n":
					if k+3 < len(strings.Split(i, "")) {
						if j+strings.Split(i, "")[k+1]+strings.Split(i, "")[k+2]+strings.Split(i, "")[k+3] == "nine" {
							line = append(line, 9)
						}
					}
				}
			}
		}
		result += line[0]*10 + line[len(line)-1]
	}
	return result
}
