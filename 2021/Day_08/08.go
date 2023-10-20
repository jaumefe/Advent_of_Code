package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	input, err := os.ReadFile("prova.txt")
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
	var led []string
	in := strings.Split(input, "\n")
	segments := map[int]int{
		0: 6, 1: 2, 2: 5, 3: 5, 4: 4, 5: 5, 6: 6, 7: 3, 8: 7, 9: 6,
	}
	for _, j := range in {
		led = append(led, strings.Trim(strings.Split(j, "|")[1], " "))
	}
	var decode = []string{"cagedb", "ab", "gcdfa", "cdfbe", "eafb", "fbcad", "cefabd", "dab", "acedgfb", "cdfgeb"}
	var sortedDecode []string
	for _, i := range decode {
		sortedDecode = append(sortedDecode, sortAlphabetically(i))
	}
	for _, i := range led {
		message := strings.Split(i, " ")
		for m, j := range message {
			if len(j) == segments[1] {
				res += 1 * int(math.Pow10(3-m))
				j = ""
			} else if len(j) == segments[4] {
				res += 4 * int(math.Pow10(3-m))
				j = ""
			} else if len(j) == segments[7] {
				res += 7 * int(math.Pow10(3-m))
				j = ""
			} else if len(j) == segments[8] {
				res += 8 * int(math.Pow10(3-m))
				j = ""
			}
			for l, k := range sortedDecode {
				if sortAlphabetically(j) == k {
					res += l * int(math.Pow10(3-m))
				}
			}
		}
		fmt.Println(res)
	}
	return res
}

func sortAlphabetically(s string) string {
	var res string
	s_temp := strings.Split(s, "")
	slices.Sort(s_temp)
	for i := 0; i < len(s_temp); i++ {
		res += s_temp[i]
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
