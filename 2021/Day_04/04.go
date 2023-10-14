package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	board  [5][5]int
	marked [5][5]bool
}

func main() {
	input, err := os.ReadFile("04.txt")
	if err != nil {
		fmt.Println("ERROR")
	}
	pt1 := part1(string(input))
	fmt.Println("The Fisrt part result is: ", pt1)
	pt2 := part2(string(input))
	fmt.Println("The Second part result is: ", pt2)
}

func part1(input string) int {
	var temp []int
	var bets []string
	var res int
	in := strings.Split(strings.Trim(input, "\n"), "\n")
	numbers := strings.Split(in[0], ",")
	numbers_int := strArrToIntArr(numbers)
	for _, j := range in[2:] {
		if j != "" {
			bets = append(bets, j)
		}
	}
	length := len(bets) / 5
	b := make([]Board, length)
	for j := 0; j < len(b); j++ {
		for i := 0; i < len(b[j].board); i++ {
			if bets[i+5*j] != "" {
				temp = strArrToIntArr(filterStrArr(strings.Split(bets[i+5*j], " ")))
				copy(b[j].board[i][:], temp)
			}
		}
	}
	for i, j := range numbers_int {
		for k := 0; k < len(b); k++ {
			check(&b[k], i, j)
			if isThereWinner(&b[k]) {
				res = j * sumMarkedBoard(&b[k])
				return res

			}
		}
	}
	return res
}

func part2(input string) int {
	var temp []int
	var bets []string
	var res int
	var winner []int
	in := strings.Split(strings.Trim(input, "\n"), "\n")
	numbers := strings.Split(in[0], ",")
	numbers_int := strArrToIntArr(numbers)
	for _, j := range in[2:] {
		if j != "" {
			bets = append(bets, j)
		}
	}
	length := len(bets) / 5
	b := make([]Board, length)
	for j := 0; j < len(b); j++ {
		for i := 0; i < len(b[j].board); i++ {
			if bets[i+5*j] != "" {
				temp = strArrToIntArr(filterStrArr(strings.Split(bets[i+5*j], " ")))
				copy(b[j].board[i][:], temp)
			}
		}
	}
	for i, j := range numbers_int {
		for k := 0; k < len(b); k++ {
			check(&b[k], i, j)
			if isThereWinner(&b[k]) {
				appendIfNotRepeated(&winner, k)
				if len(winner) == length {
					res = j * sumMarkedBoard(&b[k])
					return res
				}
			}
		}
	}
	return res
}

func appendIfNotRepeated(v *[]int, num int) {
	var temp int
	if len(*v) == 0 {
		*v = append(*v, num)
	} else {
		for _, i := range *v {
			if num != i {
				temp++
				if temp == len(*v) {
					*v = append(*v, num)
				}
			}
		}
	}
}

func check(board *Board, index int, number int) {
	for i := 0; i < len(board.board); i++ {
		for j := 0; j < len(board.board[0]); j++ {
			if number == board.board[i][j] {
				board.marked[i][j] = true
			}
		}
	}
}

func sumMarkedBoard(board *Board) int {
	var res int
	for i := 0; i < len(board.board); i++ {
		for j := 0; j < len(board.board[0]); j++ {
			if !board.marked[i][j] {
				res += board.board[i][j]
			}
		}
	}
	return res
}

func isThereWinner(board *Board) bool {
	for i := 0; i < len(board.board); i++ {
		for j := 0; j < len(board.board[0]); j++ {
			if board.marked[i][j] {
				count := 0
				// Checking a row
				for k := j; k < len(board.board[0]); k++ {
					if board.marked[i][k] {
						count++
					}
				}
				if count == 5 {
					return true
				}
			}
			if board.marked[i][j] {
				count := 0
				// Checking a column
				for k := i; k < len(board.board); k++ {
					if board.marked[k][j] {
						count++
					}
				}
				if count == 5 {
					return true
				}
			}
		}
	}
	return false
}

func strArrToIntArr(input []string) []int {
	res := make([]int, len(input))
	for i, j := range input {
		res[i], _ = strconv.Atoi(j)
	}
	return res
}

func filterStrArr(input []string) []string {
	res := make([]string, 0, 5)
	for _, j := range input {
		if j != "" {
			res = append(res, j)
		}
	}
	return res
}
