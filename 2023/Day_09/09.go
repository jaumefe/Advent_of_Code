package main

import (
	"embed"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

//go:embed inputs/*.txt
var inputFS embed.FS

func main() {
	start := time.Now()
	input1, err := inputFS.ReadFile("inputs/input.txt")
	if err != nil {
		log.Panic(err)
	}
	pt1 := part1(string(input1))
	fmt.Println("The Fisrt part result is:", pt1)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	start = time.Now()
	input2, err := inputFS.ReadFile("inputs/input.txt")
	if err != nil {
		log.Panic(err)
	}
	pt2 := part2(string(input2))
	fmt.Println("The Second part result is:", pt2)
	elapsed = time.Since(start)
	fmt.Println(elapsed)
}

func part1(input string) int {
	var result int
	in := strings.Split(input, "\n")
	var data [][]int
	for _, i := range in {
		inTmp := strings.Split(i, " ")
		data = append(data, strArrToIntArr(inTmp))
	}
	for _, i := range data {
		result += extrapolate(i) + i[len(i)-1]
	}
	return result
}

func extrapolate(data []int) int {
	var diff []int
	for i := 0; i < len(data)-1; i++ {
		diff = append(diff, data[i+1]-data[i])
	}
	equal := true
	for i := 0; i < len(diff)-1; i++ {
		if diff[i+1] != diff[i] {
			equal = false
			break
		}
		equal = true
	}
	if equal {
		return diff[len(diff)-1]
	} else {
		return extrapolate(diff) + diff[len(diff)-1]
	}
}

func extrapolatePt2(data []int) int {
	var diff []int
	for i := 0; i < len(data)-1; i++ {
		diff = append(diff, data[i+1]-data[i])
	}
	equal := true
	for i := 0; i < len(diff)-1; i++ {
		if diff[i+1] != diff[i] {
			equal = false
			break
		}
		equal = true
	}
	if equal {
		return diff[0]
	} else {
		return -extrapolatePt2(diff) + diff[0]
	}
}

func part2(input string) int {
	var result int
	in := strings.Split(input, "\n")
	var data [][]int
	for _, i := range in {
		inTmp := strings.Split(i, " ")
		data = append(data, strArrToIntArr(inTmp))
	}
	for _, i := range data {
		result += -extrapolatePt2(i) + i[0]
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
