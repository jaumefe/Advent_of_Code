package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("03.txt")
	if err != nil {
		fmt.Println("ERROR")
	}
	pt1 := part1(string(input))
	fmt.Println("The Fisrt part result is: ", pt1)
	pt2 := part2(string(input))
	fmt.Println("The Second part result is: ", pt2)
}

func part1(input string) int {
	count_zeros, count_ones := 0, 0
	gamma, epsilon := "", ""
	in := strings.Split(strings.Trim(input, "\n"), "\n")
	for j := 0; j < len(strings.Split(in[0], "")); j++ {
		for i := 0; i < len(in); i++ {
			bit := strings.Split(in[i], "")
			if bit[j] == "0" {
				count_zeros += 1
			} else {
				count_ones += 1
			}
		}
		if count_zeros > count_ones {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
		count_ones = 0
		count_zeros = 0
	}
	return strToNum(gamma) * strToNum(epsilon)
}

func part2(input string) int {
	in := strings.Split(strings.Trim(input, "\n"), "\n")
	O2, CO2 := bitCriteria(in, "1"), bitCriteria(in, "0")
	return strToNum(O2) * strToNum(CO2)
}

func bitCriteria(s []string, criteria string) string {
	count_zeros, count_ones := 0, 0
	var index_zeros []int
	var index_ones []int
	var snew []string
	for j := 0; j < len(strings.Split(s[0], "")); j++ {
		for i := 0; i < len(s); i++ {
			bit := strings.Split(s[i], "")
			if bit[j] == "0" {
				count_zeros += 1
				index_zeros = append(index_zeros, i)
			} else {
				count_ones += 1
				index_ones = append(index_ones, i)
			}
		}

		if criteria == "1" {
			if count_zeros > count_ones {
				for _, i := range index_zeros {
					snew = append(snew, s[i])
				}
			} else if count_ones > count_zeros {
				for _, i := range index_ones {
					snew = append(snew, s[i])
				}
			} else if count_ones == count_zeros {
				for _, i := range index_ones {
					snew = append(snew, s[i])
				}
			}
		} else if criteria == "0" {
			if count_zeros > count_ones {
				for _, i := range index_ones {
					snew = append(snew, s[i])
				}
			} else if count_ones > count_zeros {
				for _, i := range index_zeros {
					snew = append(snew, s[i])
				}
			} else if count_ones == count_zeros {
				for _, i := range index_zeros {
					snew = append(snew, s[i])
				}
			}
		}
		s = snew
		if len(s) == 1 {
			return s[0]
		}
		snew = nil
		count_zeros = 0
		count_ones = 0
		index_zeros = nil
		index_ones = nil
	}
	return s[0]
}

func strToNum(s string) int {
	in := strings.Split(s, "")
	res := 0
	for i := 0; i < len(in); i++ {
		val, _ := strconv.Atoi(in[i])
		res += val * int(math.Pow(2, float64(len(in)-1-i)))
	}
	return res
}
