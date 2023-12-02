package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("09.txt")
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
	var deep depth
	in := strings.Split(input, "\n")
	for _, i := range in {
		deep.depth = append(deep.depth, strArrToIntArr(strings.Split(i, "")))
	}
	deep.checkIfDeeper()
	for i, j := range deep.risk {
		for k, l := range j {
			if l {
				res += (1 + deep.depth[i][k])
			}
		}
	}
	return res
}

type depth struct {
	depth [][]int
	risk  [][]bool
	basin [][]bool
}

func part2(input string) int {
	var res int
	var deep depth
	in := strings.Split(input, "\n")
	for _, i := range in {
		deep.depth = append(deep.depth, strArrToIntArr(strings.Split(i, "")))
	}
	return res
}

func (d *depth) checkIfDeeper() [][]bool {
	// Initialisation of d.risk
	for i := 0; i < len(d.depth); i++ {
		var temp []bool
		for j := 0; j < len(d.depth[0]); j++ {
			temp = append(temp, true)
		}
		d.risk = append(d.risk, temp)
	}
	for i, j := range d.depth {
		for k, l := range j {
			if i != 0 {
				if l > d.depth[i-1][k] {
					d.risk[i][k] = false
				} else {
					d.risk[i-1][k] = false
				}
			}
			if k != 0 {
				if l > d.depth[i][k-1] {
					d.risk[i][k] = false
				} else {
					d.risk[i][k-1] = false
				}
			}
		}

	}
	return d.risk
}

func strArrToIntArr(input []string) []int {
	res := make([]int, len(input))
	for i, j := range input {
		res[i], _ = strconv.Atoi(j)
	}
	return res
}
