package main

import (
	"embed"
	"fmt"
	"log"
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

type maze struct {
	pipe     [][]string
	distance [][]int
}

func part1(input string) int {
	var result int
	in := strings.Split(input, "\n")
	var pipe maze
	for _, i := range in {
		temp := strings.Split(i, "")
		var distance []int
		if len(temp) != 0 {
			pipe.pipe = append(pipe.pipe, temp)
		}
		for j := 0; j < len(temp); j++ {
			distance = append(distance, -1)
		}
		pipe.distance = append(pipe.distance, distance)
	}
	initRow, initCol := pipe.findStart()
	var tempCoord []int
	var initCoord [][]int
	tempCoord = append(tempCoord, initRow, initCol)
	initCoord = append(initCoord, tempCoord)
	pipe.findMainLoop(initCoord, 0)
	result = pipe.findMax()
	return result
}

func (m *maze) findMax() int {
	var max int
	for i := 0; i < len(m.distance); i++ {
		for j := 0; j < len(m.distance[0]); j++ {
			if m.distance[i][j] > max {
				max = m.distance[i][j]
			}
		}
	}
	return max
}

func (m *maze) findMainLoop(coord [][]int, distance int) {
	distance += 1
	var allowed []string
	var newCoord [][]int
	for _, i := range coord {
		row, col := i[0], i[1]
		if m.pipe[row][col] != "|" {
			// Check if it can go right
			if m.pipe[row][col] != "7" && m.pipe[row][col] != "J" {
				allowed = []string{"7", "J", "-"}
				if col < len(m.distance[0])-1 {
					for _, j := range allowed {
						if m.pipe[row][col+1] == j && m.distance[row][col+1] == -1 {
							m.distance[row][col+1] = distance
							var coordTemp []int
							coordTemp = append(coordTemp, row, col+1)
							newCoord = append(newCoord, coordTemp)
							break
						}
					}
				}
			}
			// Check if it can go left
			if m.pipe[row][col] != "L" && m.pipe[row][col] != "F" {
				allowed = []string{"L", "F", "-"}
				if col > 0 {
					for _, j := range allowed {
						if m.pipe[row][col-1] == j && m.distance[row][col-1] == -1 {
							m.distance[row][col-1] = distance
							var coordTemp []int
							coordTemp = append(coordTemp, row, col-1)
							newCoord = append(newCoord, coordTemp)
							break
						}
					}
				}
			}
		}
		if m.pipe[row][col] != "-" {
			// Check if it can go up
			if m.pipe[row][col] != "7" && m.pipe[row][col] != "F" {
				allowed = []string{"7", "F", "|"}
				if row > 0 {
					for _, j := range allowed {
						if m.pipe[row-1][col] == j && m.distance[row-1][col] == -1 {
							m.distance[row-1][col] = distance
							var coordTemp []int
							coordTemp = append(coordTemp, row-1, col)
							newCoord = append(newCoord, coordTemp)
							break
						}
					}
				}
			}
			// Check if it can go down
			if m.pipe[row][col] != "L" && m.pipe[row][col] != "J" {
				allowed = []string{"L", "J", "|"}
				if row < len(m.pipe)-1 {
					for _, j := range allowed {
						if m.pipe[row+1][col] == j && m.distance[row+1][col] == -1 {
							m.distance[row+1][col] = distance
							var coordTemp []int
							coordTemp = append(coordTemp, row+1, col)
							newCoord = append(newCoord, coordTemp)
							break
						}
					}
				}
			}

		}
	}
	if len(newCoord) != 0 {
		m.findMainLoop(newCoord, distance)
	}
}

func (m *maze) findStart() (int, int) {
	var row, col int
	for i := 0; i < len(m.distance); i++ {
		for j := 0; j < len(m.distance[0]); j++ {
			if m.pipe[i][j] == "S" {
				row, col = i, j
				m.distance[i][j] = 0
				break
			}
		}
	}
	return row, col
}

func part2(input string) int {
	var result int
	in := strings.Split(input, "\n")
	var pipe maze
	for _, i := range in {
		temp := strings.Split(i, "")
		var distance []int
		if len(temp) != 0 {
			pipe.pipe = append(pipe.pipe, temp)
		}
		for j := 0; j < len(temp); j++ {
			distance = append(distance, -1)
		}
		pipe.distance = append(pipe.distance, distance)
	}
	initRow, initCol := pipe.findStart()
	var tempCoord []int
	var initCoord [][]int
	tempCoord = append(tempCoord, initRow, initCol)
	initCoord = append(initCoord, tempCoord)
	vertex := pipe.findMainLoopPt2(initCoord, 0)
	var first, second [][]int
	for i := 0; i < len(vertex)-1; i += 2 {
		first = append(first, vertex[i])
	}
	for i := 1; i < len(vertex)-1; i += 2 {
		second = append(second, vertex[i])
	}
	var orderedVertex [][]int
	orderedVertex = append(orderedVertex, initCoord[0])
	for i := 0; i < len(first); i++ {
		orderedVertex = append(orderedVertex, first[i])
	}
	for i := 0; i < len(second); i++ {
		orderedVertex = append(orderedVertex, second[len(second)-1-i])
	}
	twoArea, I := shoelaceFormula(orderedVertex)
	if twoArea < 0 {
		twoArea = -twoArea
	}
	result = int((twoArea - (I + 1) + 3) / 2)
	return result
}

func (m *maze) findMainLoopPt2(coord [][]int, distance int) [][]int {
	distance += 1
	var allowed []string
	var newCoord [][]int
	var temp [][]int
	for _, i := range coord {
		row, col := i[0], i[1]
		if m.pipe[row][col] != "|" {
			// Check if it can go right
			if m.pipe[row][col] != "7" && m.pipe[row][col] != "J" {
				allowed = []string{"7", "J", "-"}
				if col < len(m.distance[0])-1 {
					for _, j := range allowed {
						if m.pipe[row][col+1] == j && m.distance[row][col+1] == -1 {
							m.distance[row][col+1] = distance
							var coordTemp []int
							coordTemp = append(coordTemp, row, col+1)
							newCoord = append(newCoord, coordTemp)
							break
						}
					}
				}
			}
			// Check if it can go left
			if m.pipe[row][col] != "L" && m.pipe[row][col] != "F" {
				allowed = []string{"L", "F", "-"}
				if col > 0 {
					for _, j := range allowed {
						if m.pipe[row][col-1] == j && m.distance[row][col-1] == -1 {
							m.distance[row][col-1] = distance
							var coordTemp []int
							coordTemp = append(coordTemp, row, col-1)
							newCoord = append(newCoord, coordTemp)
							break
						}
					}
				}
			}
		}
		if m.pipe[row][col] != "-" {
			// Check if it can go up
			if m.pipe[row][col] != "7" && m.pipe[row][col] != "F" {
				allowed = []string{"7", "F", "|"}
				if row > 0 {
					for _, j := range allowed {
						if m.pipe[row-1][col] == j && m.distance[row-1][col] == -1 {
							m.distance[row-1][col] = distance
							var coordTemp []int
							coordTemp = append(coordTemp, row-1, col)
							newCoord = append(newCoord, coordTemp)
							break
						}
					}
				}
			}
			// Check if it can go down
			if m.pipe[row][col] != "L" && m.pipe[row][col] != "J" {
				allowed = []string{"L", "J", "|"}
				if row < len(m.pipe)-1 {
					for _, j := range allowed {
						if m.pipe[row+1][col] == j && m.distance[row+1][col] == -1 {
							m.distance[row+1][col] = distance
							var coordTemp []int
							coordTemp = append(coordTemp, row+1, col)
							newCoord = append(newCoord, coordTemp)
							break
						}
					}
				}
			}

		}
	}
	if len(newCoord) != 0 {
		temp = m.findMainLoopPt2(newCoord, distance)
	}
	for i := 0; i < len(temp); i++ {
		newCoord = append(newCoord, temp[i])
	}
	return newCoord
}

func shoelaceFormula(vertex [][]int) (int, int) {
	var result int
	var pick int
	for i := 0; i < len(vertex); i++ {
		var x2, y2 int
		x1, y1 := vertex[i][0], vertex[i][1]
		if i+1 >= len(vertex) {
			x2, y2 = vertex[0][0], vertex[0][1]
		} else {
			x2, y2 = vertex[i+1][0], vertex[i+1][1]
		}
		result += x1*y2 - x2*y1
		pick++
	}
	return result, pick
}
