package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	// origin x1, y1 values
	origin [][]int
	// dest x2, y2 values
	dest [][]int
}

func main() {
	input, err := os.ReadFile("05.txt")
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
	var coord Coord
	in := strings.Split(input, "\n")
	for _, i := range in {
		in_temp := strings.Split(i, "->")
		for j, k := range in_temp {
			coord_temp := strArrToIntArr(strings.Split(strings.Trim(k, " "), ","))
			if j%2 == 0 {
				coord.origin = append(coord.origin, coord_temp)
			} else {
				coord.dest = append(coord.dest, coord_temp)
			}
		}
	}
	row, col := findMaxRowAndCol(coord)
	max := int(math.Max(float64(row), float64(col)))
	grid := make([][]int, max+1)

	for i := 0; i <= max; i++ {
		for j := 0; j <= max; j++ {
			grid[i] = append(grid[i], 0)
		}
	}
	for i := 0; i < len(coord.origin); i++ {
		if isTheSameColumn(coord.origin[i], coord.dest[i]) {
			if coord.origin[i][0] < coord.dest[i][0] {
				for j := coord.origin[i][0]; j <= coord.dest[i][0]; j++ {
					grid[coord.dest[i][1]][j]++
				}
			} else {
				for j := coord.origin[i][0]; j >= coord.dest[i][0]; j-- {
					grid[coord.dest[i][1]][j]++
				}
			}
		}
		if isTheSameRow(coord.origin[i], coord.dest[i]) {
			if coord.origin[i][1] < coord.dest[i][1] {
				for j := coord.origin[i][1]; j <= coord.dest[i][1]; j++ {
					grid[j][coord.dest[i][0]]++
				}
			} else {
				for j := coord.origin[i][1]; j >= coord.dest[i][1]; j-- {
					grid[j][coord.dest[i][0]]++
				}
			}
		}
	}
	for _, i := range grid {
		for _, j := range i {
			if j >= 2 {
				res++
			}
		}
	}
	return res
}

func part2(input string) int {
	var res int
	var coord Coord
	in := strings.Split(input, "\n")
	for _, i := range in {
		in_temp := strings.Split(i, "->")
		for j, k := range in_temp {
			coord_temp := strArrToIntArr(strings.Split(strings.Trim(k, " "), ","))
			if j%2 == 0 {
				coord.origin = append(coord.origin, coord_temp)
			} else {
				coord.dest = append(coord.dest, coord_temp)
			}
		}
	}
	row, col := findMaxRowAndCol(coord)
	max := int(math.Max(float64(row), float64(col)))
	grid := make([][]int, max+1)

	for i := 0; i <= max; i++ {
		for j := 0; j <= max; j++ {
			grid[i] = append(grid[i], 0)
		}
	}
	for i := 0; i < len(coord.origin); i++ {
		// Check if the two points are on the same column and if true, move in the row
		if isTheSameColumn(coord.origin[i], coord.dest[i]) {
			if coord.origin[i][0] < coord.dest[i][0] {
				for j := coord.origin[i][0]; j <= coord.dest[i][0]; j++ {
					grid[coord.dest[i][1]][j]++
				}
			} else {
				for j := coord.origin[i][0]; j >= coord.dest[i][0]; j-- {
					grid[coord.dest[i][1]][j]++
				}
			}
			// Check if the two points are on the same row and if true, move in the column
		} else if isTheSameRow(coord.origin[i], coord.dest[i]) {
			if coord.origin[i][1] < coord.dest[i][1] {
				for j := coord.origin[i][1]; j <= coord.dest[i][1]; j++ {
					grid[j][coord.dest[i][0]]++
				}
			} else {
				for j := coord.origin[i][1]; j >= coord.dest[i][1]; j-- {
					grid[j][coord.dest[i][0]]++
				}
			}
			// Move in diagonal
		} else {
			if (coord.origin[i][0] < coord.dest[i][0]) && (coord.origin[i][1] < coord.dest[i][1]) {
				for j, k := coord.origin[i][0], coord.origin[i][1]; j <= coord.dest[i][0] && k <= coord.dest[i][1]; j, k = j+1, k+1 {
					grid[k][j]++
				}
			}
			if (coord.origin[i][0] > coord.dest[i][0]) && (coord.origin[i][1] > coord.dest[i][1]) {
				for j, k := coord.origin[i][0], coord.origin[i][1]; j >= coord.dest[i][0] && k >= coord.dest[i][1]; j, k = j-1, k-1 {
					grid[k][j]++
				}
			}
			if (coord.origin[i][0] < coord.dest[i][0]) && (coord.origin[i][1] > coord.dest[i][1]) {
				for j, k := coord.origin[i][0], coord.origin[i][1]; j <= coord.dest[i][0] && k >= coord.dest[i][1]; j, k = j+1, k-1 {
					grid[k][j]++
				}
			}
			if (coord.origin[i][0] > coord.dest[i][0]) && (coord.origin[i][1] < coord.dest[i][1]) {
				for j, k := coord.origin[i][0], coord.origin[i][1]; j >= coord.dest[i][0] && k <= coord.dest[i][1]; j, k = j-1, k+1 {
					grid[k][j]++
				}
			}
		}
	}

	for _, i := range grid {
		for _, j := range i {
			if j >= 2 {
				res++
			}
		}
	}
	return res
}

func isTheSameRow(origin []int, dest []int) bool {
	return origin[0] == dest[0]
}

func isTheSameColumn(origin []int, dest []int) bool {
	return origin[1] == dest[1]
}

func findMaxRowAndCol(coord Coord) (int, int) {
	var origin_x []int
	var origin_y []int
	var dest_x []int
	var dest_y []int
	var max_x int
	var max_y int
	for _, i := range coord.origin {
		origin_x = append(origin_x, i[0])
		origin_y = append(origin_y, i[1])
	}
	for _, i := range coord.dest {
		dest_x = append(dest_x, i[0])
		dest_y = append(dest_y, i[1])
	}
	for i, j := range origin_x {
		max_x = int(math.Max(float64(j), float64(max_x)))
		max_y = int(math.Max(float64(origin_y[i]), float64(max_y)))
	}
	for i, j := range dest_x {
		max_x = int(math.Max(float64(j), float64(max_x)))
		max_y = int(math.Max(float64(dest_y[i]), float64(max_y)))
	}
	return max_x, max_y
}

func strArrToIntArr(input []string) []int {
	res := make([]int, len(input))
	for i, j := range input {
		res[i], _ = strconv.Atoi(j)
	}
	return res
}
