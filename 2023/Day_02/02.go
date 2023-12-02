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
	var sets []string
	for _, i := range in {
		sets = append(sets, strings.Split(i, ": ")[1])
	}
	var i_temp []string
	for index, i := range sets {
		i_temp = strings.Split(strings.Replace(i, ",", "", -1), "; ")
		colors := map[string]int{"red": 0, "blue": 0, "green": 0}
		isPossible := true
		for _, k := range i_temp {
			var set []string

			for _, j := range strings.Split(k, " ") {
				set = append(set, j)
			}
			for j := 1; j < len(set); j += 2 {
				colorsNum, _ := strconv.Atoi(set[j-1])
				colors[set[j]] = colorsNum
				if (colors["blue"] > 14) || (colors["green"] > 13) || (colors["red"] > 12) {
					isPossible = false
				}
			}
		}
		if isPossible {
			result += index + 1
		}
	}

	return result
}

func part2(input string) int {
	var result int
	in := strings.Split(input, "\n")
	var sets []string
	for _, i := range in {
		sets = append(sets, strings.Split(i, ": ")[1])
	}
	var i_temp []string
	for _, i := range sets {
		i_temp = strings.Split(strings.Replace(i, ",", "", -1), "; ")
		colors := map[string]int{"red": 0, "blue": 0, "green": 0}
		for _, k := range i_temp {
			var set []string

			for _, j := range strings.Split(k, " ") {
				set = append(set, j)
			}
			for j := 1; j < len(set); j += 2 {
				colorsNum, _ := strconv.Atoi(set[j-1])
				if colors[set[j]] <= colorsNum {
					colors[set[j]] = colorsNum
				}
			}
		}
		//fmt.Println(colors)
		result += colors["blue"] * colors["red"] * colors["green"]
	}
	return result
}
