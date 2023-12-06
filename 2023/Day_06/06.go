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
	var time, distance []string
	for index, i := range strings.Split(input, "\n") {
		for subindex, j := range strings.Split(strings.Join(strings.Fields(i), " "), " ") {
			if index == 0 && subindex > 0 {
				time = append(time, j)
			} else if index == 1 && subindex > 0 {
				distance = append(distance, j)
			}
		}
	}
	timeInt := strArrToIntArr(time)
	distanceInt := strArrToIntArr(distance)
	result = 1
	for i, t := range timeInt {
		var count int
		t_temp := int(t / 2)
		for j := 0; j <= t_temp; j++ {
			speed := 1 * j
			disTravel := speed * (t - j)
			if disTravel > distanceInt[i] {
				count++
			}
		}
		if t%2 != 0 {
			result *= count * 2
		} else {
			result *= count*2 - 1
		}

	}

	return result
}

func part2(input string) int {
	var result int
	var time, distance []string
	for index, i := range strings.Split(input, "\n") {
		for subindex, j := range strings.Split(strings.Join(strings.Fields(i), " "), " ") {
			if index == 0 && subindex > 0 {
				time = append(time, j)
			} else if index == 1 && subindex > 0 {
				distance = append(distance, j)
			}
		}
	}

	var distanceInt, timeInt int
	var timeTemp, distTemp string
	for i := 0; i < len(time); i++ {
		timeTemp += time[i]
		distTemp += distance[i]
	}
	timeInt, _ = strconv.Atoi(timeTemp)
	distanceInt, _ = strconv.Atoi(distTemp)
	var count int
	result = 1
	for i := 0; i <= int(timeInt/2); i++ {

		speed := 1 * i
		disTravel := speed * (timeInt - i)
		if disTravel > distanceInt {
			count++
		}
	}
	if timeInt%2 != 0 {
		result *= count * 2
	} else {
		result *= count*2 - 1
	}

	return result
}

func strArrToIntArr(input []string) []int {
	res := make([]int, len(input))
	for i, j := range input {
		res[i], _ = strconv.Atoi(j)
	}
	return res
}
