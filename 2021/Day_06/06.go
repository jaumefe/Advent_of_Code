package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("06.txt")
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
	var days int
	in := strings.Split(input, ",")
	lanternFish := strArrToIntArr(in)
	for days < 80 {
		count := 0
		for i, j := range lanternFish {
			if j != 0 {
				lanternFish[i]--
			} else {
				lanternFish[i] = 6
				count++
			}
		}
		for i := 0; i < count; i++ {
			lanternFish = append(lanternFish, 8)
		}
		days++
	}
	res = len(lanternFish)
	return res
}

// After inspiration reading on Internet, here it is a better solution
// It is not needed to create a vector with all the elements
// A better approach is to count how many fishes are on each state
// We create a initial map up to 9 elements (from 0 to 8, corresponding to the number of days before breeding)
// After each day passed we follow this technique:
// map[0] -> new_offspring
// map[1] -> map[0]; map[2] -> map[1]; ....; map[8] -> map[7]
// new_offspring -> map[8]
// As new offspring appears is important to take into account that there has to be added the same number to map[6], because parents in 6 days will be ready to breed again

func part2(input string) int {
	var res int
	var days int
	in := strArrToIntArr(strings.Split(input, ","))
	lanternFish := make(map[int]int)
	for i := 0; i < 9; i++ {
		lanternFish[i] = 0
	}
	for _, j := range in {
		lanternFish[j]++
	}

	for days < 256 {
		offspring := lanternFish[0]
		for i := 1; i < 9; i++ {
			lanternFish[i-1] = lanternFish[i]
		}
		lanternFish[8] = offspring
		lanternFish[6] += lanternFish[8]
		days++
	}
	for i := 0; i < 9; i++ {
		res += lanternFish[i]
	}
	return res
}

func strArrToIntArr(input []string) []int {
	res := make([]int, len(input))
	for i, j := range input {
		res[i], _ = strconv.Atoi(j)
	}
	return res
}
