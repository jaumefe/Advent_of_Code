package main

import (
	"embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed inputs/*.txt
var inputFS embed.FS

type partNumber struct {
	symbol [][]bool
	number [][]bool
	data   [][]string
}

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

func part1(input string) int {
	var result int
	var pn partNumber
	in := strings.Split(input, "\n")
	for _, i := range in {
		pn.data = append(pn.data, strings.Split(i, ""))
	}
	for _, i := range pn.data {
		var symbol []bool
		var num []bool
		for _, j := range i {
			if j != "." {
				_, err := strconv.Atoi(j)
				if err != nil {
					symbol = append(symbol, true)
					num = append(num, false)
				} else {
					symbol = append(symbol, false)
					num = append(num, true)
				}
			} else {
				symbol = append(symbol, false)
				num = append(num, false)
			}
		}
		pn.symbol = append(pn.symbol, symbol)
		pn.number = append(pn.number, num)
	}
	result = pn.sumDataNextToSymbol()
	return result
}

func part2(input string) int {
	var result int
	var pn partNumber
	in := strings.Split(input, "\n")
	for _, i := range in {
		pn.data = append(pn.data, strings.Split(i, ""))
	}
	for _, i := range pn.data {
		var symbol []bool
		var num []bool
		for _, j := range i {
			if j == "*" {
				symbol = append(symbol, true)
				num = append(num, false)
			} else {
				symbol = append(symbol, false)
				_, err := strconv.Atoi(j)
				if err == nil {
					num = append(num, true)
				} else {
					num = append(num, false)
				}
			}
		}
		pn.symbol = append(pn.symbol, symbol)
		pn.number = append(pn.number, num)
	}
	result = pn.findGearRatio()
	return result
}

func (pn *partNumber) findGearRatio() int {
	var result int
	for row, i := range pn.symbol {
		for col, j := range i {
			result_temp := 1
			if j {
				if pn.partNumberSurrounding(row, col) == 2 {
					if row > 0 {
						if pn.number[row-1][col] {
							result_temp *= pn.findFullNumber(row-1, col)
						}
						if col > 0 {
							if pn.number[row-1][col-1] {
								result_temp *= pn.findFullNumber(row-1, col-1)
							}
						}
						if col < len(pn.number[0]) {
							if pn.number[row-1][col+1] {
								result_temp *= pn.findFullNumber(row-1, col+1)
							}
						}
					}
					if row < len(pn.number) {
						if pn.number[row+1][col] {
							result_temp *= pn.findFullNumber(row+1, col)
						}
						if col > 0 {
							if pn.number[row+1][col-1] {
								result_temp *= pn.findFullNumber(row+1, col-1)
							}
						}
						if col < len(pn.number[0]) {
							if pn.number[row+1][col+1] {
								result_temp *= pn.findFullNumber(row+1, col+1)
							}
						}
					}
					if col > 0 {
						if pn.number[row][col-1] {
							result_temp *= pn.findFullNumber(row, col-1)
						}
					}
					if col < len(pn.number[0]) {
						if pn.number[row][col+1] {
							result_temp *= pn.findFullNumber(row, col+1)
						}
					}
				}
			}
			if result_temp != 1 {
				result += result_temp
			}
		}
	}
	return result
}

func (pn *partNumber) partNumberSurrounding(row int, col int) int {
	var result int
	if row > 0 {
		if !pn.number[row-1][col] {
			if pn.number[row-1][col-1] {
				result++
			}
			if pn.number[row-1][col+1] {
				result++
			}
		} else {
			result++
		}
	}
	if row < len(pn.number) {
		if !pn.number[row+1][col] {
			if pn.number[row+1][col-1] {
				result++
			}
			if pn.number[row+1][col+1] {
				result++
			}
		} else {
			result++
		}
	}
	if col > 0 {
		if pn.number[row][col-1] {
			result++
		}
	}
	if col < len(pn.number[0]) {
		if pn.number[row][col+1] {
			result++
		}
	}
	return result
}

func (pn *partNumber) sumDataNextToSymbol() int {
	var result int
	for row, i := range pn.symbol {
		for col, j := range i {
			if j {
				if row > 0 {
					if pn.number[row-1][col] {
						result += pn.findFullNumber(row-1, col)
					}
					if col > 0 {
						if pn.number[row-1][col-1] {
							result += pn.findFullNumber(row-1, col-1)
						}
					}
					if col < len(pn.number[0]) {
						if pn.number[row-1][col+1] {
							result += pn.findFullNumber(row-1, col+1)
						}
					}
				}
				if row < len(pn.number) {
					if pn.number[row+1][col] {
						result += pn.findFullNumber(row+1, col)
					}
					if col > 0 {
						if pn.number[row+1][col-1] {
							result += pn.findFullNumber(row+1, col-1)
						}
					}
					if col < len(pn.number[0]) {
						if pn.number[row+1][col+1] {
							result += pn.findFullNumber(row+1, col+1)
						}
					}
				}
				if col > 0 {
					if pn.number[row][col-1] {
						result += pn.findFullNumber(row, col-1)
					}
				}
				if col < len(pn.number[0]) {
					if pn.number[row][col+1] {
						result += pn.findFullNumber(row, col+1)
					}
				}
			}
		}
	}
	return result
}

func (pn *partNumber) findFullNumber(row int, col int) int {
	var result int
	var num string
	for i := col; i > -1; i-- {
		if !pn.number[row][i] {
			break
		}
		num = pn.data[row][i] + num
		pn.number[row][i] = false
	}
	for i := col + 1; i < len(pn.number); i++ {
		if !pn.number[row][i] {
			break
		}
		num += pn.data[row][i]
		pn.number[row][i] = false
	}
	result, _ = strconv.Atoi(num)
	return result
}
