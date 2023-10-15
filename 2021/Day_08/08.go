package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("08.txt")
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
	var led []string
	segments := map[int]int{
		0: 6, 1: 2, 2: 5, 3: 5, 4: 4, 5: 5, 6: 6, 7: 3, 8: 7, 9: 6,
	}
	in := strings.Split(input, "\n")
	for _, j := range in {
		led = append(led, strings.Trim(strings.Split(j, "|")[1], " "))
	}
	for _, j := range led {
		message := strings.Split(j, " ")
		for _, i := range message {
			if len(i) == segments[1] || len(i) == segments[4] || len(i) == segments[7] || len(i) == segments[8] {
				res++
			}
		}
	}
	return res
}

func part2(input string) int {
	var res int
	return res
}

func strArrToIntArr(input []string) []int {
	res := make([]int, len(input))
	for i, j := range input {
		res[i], _ = strconv.Atoi(j)
	}
	return res
}
