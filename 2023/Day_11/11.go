package main

import (
	"embed"
	"fmt"
	"log"
	"math"
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
	fmt.Println("The First part result is:", pt1)
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

type space struct {
	galaxy         [][]bool
	coord          [][]int
	expandedCoords [][]int
}

func part1(input string) int {
	var result int
	var universe space
	in := strings.Split(input, "\n")
	for idx, i := range in {
		var row []bool
		for idy, j := range strings.Split(i, "") {
			if j == "#" {
				var coord []int
				coord = append(coord, idx, idy)
				universe.coord = append(universe.coord, coord)
				row = append(row, true)
			} else {
				row = append(row, false)
			}
		}
		universe.galaxy = append(universe.galaxy, row)
	}
	universe.expandedUniverse()
	result = universe.findDistance()
	return result
}

func (u *space) expandedUniverse() {
	// Find empty columns
	cols := make([]bool, len(u.galaxy[0]))
	for _, i := range u.coord {
		cols[i[1]] = true
	}
	var emptyCols []int
	for index, i := range cols {
		if !i {
			emptyCols = append(emptyCols, index)
		}
	}
	// Find empty rows
	rows := make([]bool, len(u.galaxy))
	for _, i := range u.coord {
		rows[i[0]] = true
	}
	var emptyRows []int
	for index, i := range rows {
		if !i {
			emptyRows = append(emptyRows, index)
		}
	}
	// Compute New Coordinates
	for _, i := range u.coord {
		var tempCoord []int
		var rowAdd int
		for j := 0; j < len(emptyRows); j++ {
			if i[0] < emptyRows[j] {
				break
			} else {
				rowAdd += 1
			}
		}
		var colAdd int
		for j := 0; j < len(emptyCols); j++ {
			if i[1] < emptyCols[j] {
				break
			} else {
				colAdd += 1
			}
		}
		tempCoord = append(tempCoord, i[0]+rowAdd, i[1]+colAdd)
		u.expandedCoords = append(u.expandedCoords, tempCoord)
	}

}

func (u *space) findDistance() int {
	var result int
	for i := 0; i < len(u.expandedCoords); i++ {
		for j := i; j < len(u.expandedCoords); j++ {
			xDiff := math.Abs(float64(u.expandedCoords[i][0] - u.expandedCoords[j][0]))
			yDiff := math.Abs(float64(u.expandedCoords[i][1] - u.expandedCoords[j][1]))
			result += int(xDiff) + int(yDiff)
		}
	}
	return result
}

func part2(input string) int {
	var result int
	var universe space
	in := strings.Split(input, "\n")
	for idx, i := range in {
		var row []bool
		for idy, j := range strings.Split(i, "") {
			if j == "#" {
				var coord []int
				coord = append(coord, idx, idy)
				universe.coord = append(universe.coord, coord)
				row = append(row, true)
			} else {
				row = append(row, false)
			}
		}
		universe.galaxy = append(universe.galaxy, row)
	}
	universe.expandedUniversePt2()
	result = universe.findDistance()
	return result
}

func (u *space) expandedUniversePt2() {
	// Find empty columns
	cols := make([]bool, len(u.galaxy[0]))
	for _, i := range u.coord {
		cols[i[1]] = true
	}
	var emptyCols []int
	for index, i := range cols {
		if !i {
			emptyCols = append(emptyCols, index)
		}
	}
	// Find empty rows
	rows := make([]bool, len(u.galaxy))
	for _, i := range u.coord {
		rows[i[0]] = true
	}
	var emptyRows []int
	for index, i := range rows {
		if !i {
			emptyRows = append(emptyRows, index)
		}
	}
	// Compute New Coordinates
	for _, i := range u.coord {
		var tempCoord []int
		var rowAdd int
		for j := 0; j < len(emptyRows); j++ {
			if i[0] < emptyRows[j] {
				break
			} else {
				rowAdd += 1000000 - 1
			}
		}
		var colAdd int
		for j := 0; j < len(emptyCols); j++ {
			if i[1] < emptyCols[j] {
				break
			} else {
				colAdd += 1000000 - 1
			}
		}
		tempCoord = append(tempCoord, i[0]+rowAdd, i[1]+colAdd)
		u.expandedCoords = append(u.expandedCoords, tempCoord)
	}

}
