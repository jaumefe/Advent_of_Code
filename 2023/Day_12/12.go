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
	input1, err := inputFS.ReadFile("inputs/test.txt")
	if err != nil {
		log.Panic(err)
	}
	pt1 := part1(string(input1))
	fmt.Println("The First part result is:", pt1)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	start = time.Now()
	input2, err := inputFS.ReadFile("inputs/test.txt")
	if err != nil {
		log.Panic(err)
	}
	pt2 := part2(string(input2))
	fmt.Println("The Second part result is:", pt2)
	elapsed = time.Since(start)
	fmt.Println(elapsed)
}

type springs struct {
	dist    []string
	distNum [][]int
}

func part1(input string) int {
	var result int
	var springs springs
	in := strings.Split(input, "\n")
	springs.distNum = make([][]int, len(in))
	for idx, i := range in {
		springs.dist = append(springs.dist, strings.Split(i, " ")[0])
		springs.distNum[idx] = strArrToIntArr(strings.Split(strings.Split(i, " ")[1], ","))
	}
	fmt.Println(springs.distNum)

	return result
}

func part2(input string) int {
	var result int
	return result
}

func strArrToIntArr(input []string) []int {
	res := make([]int, len(input))
	for i, j := range input {
		res[i], _ = strconv.Atoi(j)
	}
	return res
}
