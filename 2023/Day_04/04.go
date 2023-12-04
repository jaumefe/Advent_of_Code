package main

import (
	"embed"
	"fmt"
	"log"
	"math"
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

type Scrathcard struct {
	winner  [][]string
	numbers [][]string
	copy    map[int]int
}

func part1(input string) int {
	var result int
	var sc Scrathcard
	in := strings.Split(strings.Replace(input, "  ", " ", -1), "\n")
	var card []string
	for _, i := range in {
		card = append(card, strings.Split(i, ": ")[1])
	}
	var winning []string
	var num []string
	for _, i := range card {
		winning = append(winning, strings.Split(i, " | ")[0])
		num = append(num, strings.Split(i, " | ")[1])
	}
	for i := 0; i < len(winning); i++ {
		sc.winner = append(sc.winner, strings.Split(winning[i], " "))
	}
	for i := 0; i < len(num); i++ {
		sc.numbers = append(sc.numbers, strings.Split(num[i], " "))
	}
	result = sc.pointSum()
	return result
}

func (sc *Scrathcard) pointSum() int {
	var result int
	var count int
	for index, i := range sc.numbers {
		for _, j := range i {
			for _, k := range sc.winner[index] {
				if j == k {
					count++
				}
			}
		}
		if count > 0 {
			result += int(math.Pow(2, float64(count-1)))
			count = 0
		}
	}
	return result
}

func part2(input string) int {
	var result int
	var sc Scrathcard
	in := strings.Split(strings.Replace(input, "  ", " ", -1), "\n")
	var card []string
	for _, i := range in {
		card = append(card, strings.Split(i, ": ")[1])
	}
	var winning []string
	var num []string
	for _, i := range card {
		winning = append(winning, strings.Split(i, " | ")[0])
		num = append(num, strings.Split(i, " | ")[1])
	}
	for i := 0; i < len(winning); i++ {
		sc.winner = append(sc.winner, strings.Split(winning[i], " "))
	}
	for i := 0; i < len(num); i++ {
		sc.numbers = append(sc.numbers, strings.Split(num[i], " "))
	}
	sc.copy = make(map[int]int)
	for i := 0; i < len(num); i++ {
		sc.copy[i] = 1
	}
	sc.scrathCardCopyFinder()
	for i := 0; i < len(sc.copy); i++ {
		result += sc.copy[i]
	}
	return result
}

func (sc *Scrathcard) scrathCardCopyFinder() {
	var count int
	for index, i := range sc.numbers {
		for _, j := range i {
			for _, k := range sc.winner[index] {
				if j == k {
					count++
				}
			}
		}
		for j := 1 + index; j < 1+index+count; j++ {
			sc.copy[j] += sc.copy[index]
		}
		count = 0
	}
}
