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

type node struct {
	left  string
	right string
	name  string
}

func part1(input string) int {
	var result int
	var nodes []node
	in := strings.Split(input, "\n")
	direction := in[0]
	for _, i := range in[2:] {
		var nodeTemp node
		i = strings.Replace(strings.Replace(i, "(", "", -1), ")", "", -1)
		nodeTemp.name = string(strings.Split(i, " = ")[0])
		leftTemp, rightTemp := strings.Split((strings.Split(i, " = ")[1]), ", ")[0], strings.Split((strings.Split(i, " = ")[1]), ", ")[1]
		nodeTemp.left, nodeTemp.right = leftTemp, rightTemp
		nodes = append(nodes, nodeTemp)
	}
	currentNode := "AAA"
	var count, index int
	for currentNode != "ZZZ" {
		if index == len(direction) {
			index = 0
		}
		dir := string(direction[index])
		var node node
		for _, i := range nodes {
			if i.name == currentNode {
				node = i
				break
			}
		}
		switch dir {
		case "L":
			currentNode = node.left
		case "R":
			currentNode = node.right
		}
		count++
		index++
	}
	result = count
	return result
}

func part2(input string) int {
	var result int
	var nodes []node
	in := strings.Split(input, "\n")
	direction := in[0]
	for _, i := range in[2:] {
		var nodeTemp node
		i = strings.Replace(strings.Replace(i, "(", "", -1), ")", "", -1)
		nodeTemp.name = string(strings.Split(i, " = ")[0])
		leftTemp, rightTemp := strings.Split((strings.Split(i, " = ")[1]), ", ")[0], strings.Split((strings.Split(i, " = ")[1]), ", ")[1]
		nodeTemp.left, nodeTemp.right = leftTemp, rightTemp
		nodes = append(nodes, nodeTemp)
	}
	var total []int
	for _, n := range nodes {
		var count, index int
		node := n
		if string(n.name[2]) == "A" {
			currentNode := "A"
			for currentNode != "Z" {
				if index == len(direction) {
					index = 0
				}
				dir := string(direction[index])
				switch dir {
				case "L":
					currentNode = string(node.left[2])
					for _, i := range nodes {
						if i.name == node.left {
							node = i
							break
						}
					}
				case "R":
					currentNode = string(node.right[2])
					for _, i := range nodes {
						if i.name == node.right {
							node = i
							break
						}
					}
				}
				count++
				index++
			}

		}
		if count != 0 {
			total = append(total, count)
		}
	}
	first := total[0]
	for i := 1; i < len(total); i++ {
		second := total[i]
		first = LCM(first, second)
	}
	result = first
	return result
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
