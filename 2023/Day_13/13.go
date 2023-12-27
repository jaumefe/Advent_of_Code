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

func part1(input string) int {
	var result int
	in := strings.Split(input, "\n")
	var mirror []string
	for idx, i := range in {
		if len(i) != 0 {
			mirror = append(mirror, i)
		}
		if (len(i) == 0) || (idx == len(in)-1) {
			result += findHorizontalAxe(mirror)*100 + findVerticalAxe(mirror)
			mirror = make([]string, 0)
		}
	}
	return result
}

func findHorizontalAxe(data []string) int {
	result := 0
	for i := 0; i < len(data)-1; i++ {
		result++
		if data[i] == data[i+1] {
			var symetric bool
			lim := int(math.Min(float64(result), float64(len(data)-result)))
			for j := 0; j < lim; j++ {
				symetric = true
				if data[i-j] != data[i+1+j] {
					symetric = false
					break
				}
			}
			if symetric {
				return result
			}
		}
	}
	return 0
}

func findVerticalAxe(data []string) int {
	// Rearrange columns into rows
	newData := make([]string, len(data[0]))
	for j := 0; j < len(data[0]); j++ {
		for i := 0; i < len(data); i++ {
			newData[j] += string(data[i][j])
		}
	}
	result := findHorizontalAxe(newData)
	return result
}

type mirror struct {
	data []string
}

func part2(input string) int {
	var result int
	in := strings.Split(input, "\n")
	var mirror mirror
	for idx, i := range in {
		if len(i) != 0 {
			mirror.data = append(mirror.data, i)
		}
		if (len(i) == 0) || (idx == len(in)-1) {
			result += mirror.findHorizontalAxePt2() * 100
			result += mirror.findVerticalAxePt2()
			mirror.data = make([]string, 0)
		}
	}
	return result
}

func checkIfSmudge(row1 string, row2 string) bool {
	var count int
	//var index int
	for i := 0; i < len(row1); i++ {
		if row1[i] != row2[i] {
			count++
			//index = i
		}
		if count > 1 {
			return false
		}
	}
	/*var rowTemp string
	if string(row1[index]) == "." {
		for i := 0; i < len(row1); i++ {
			if i != index {
				rowTemp += string(row1[i])
			} else {
				rowTemp += "#"
			}
		}
	} else if string(row2[index]) == "." {
		for i := 0; i < len(row1); i++ {
			if i != index {
				rowTemp += string(row1[i])
			} else {
				rowTemp += "#"
			}
		}
	}*/
	return true
}

func (m *mirror) findHorizontalAxePt2() int {
	result := 0
	var smudge bool
	for i := 0; i < len(m.data)-1; i++ {
		result++
		if m.data[i] != m.data[i+1] {
			smudge = checkIfSmudge(m.data[i], m.data[i+1])
			if smudge {
				//m.data[i+1] = m.data[i]
			}
		}
		if m.data[i] == m.data[i+1] || smudge {
			var symetric bool
			lim := int(math.Min(float64(result), float64(len(m.data)-result)))
			for j := 0; j < lim; j++ {
				symetric = true
				if m.data[i-j] != m.data[i+1+j] {
					symetric = checkIfSmudge(m.data[i-j], m.data[i+1+j])
					if symetric {
						//m.data[i-j] = m.data[i+1+j]
					} else {
						break
					}
				}
			}
			if symetric {
				return result
			}
		}
	}
	return 0
}

func (m *mirror) findVerticalAxePt2() int {
	// Rearrange columns into rows
	var newData mirror
	newData.data = make([]string, len(m.data[0]))
	for j := 0; j < len(m.data[0]); j++ {
		for i := 0; i < len(m.data); i++ {
			newData.data[j] += string(m.data[i][j])
		}
	}
	result := newData.findHorizontalAxePt2()
	return result
}
