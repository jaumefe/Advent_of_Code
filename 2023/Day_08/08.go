package main

import (
	"embed"
	"fmt"
	"log"
	"strings"
)

//go:embed inputs/*.txt
var inputFS embed.FS

func main() {
	input1, err := inputFS.ReadFile("inputs/input.txt")
	if err != nil {
		log.Panic(err)
	}
	pt1 := part1(string(input1))
	fmt.Println("The Fisrt part result is:", pt1)
	input2, err := inputFS.ReadFile("inputs/test.txt")
	if err != nil {
		log.Panic(err)
	}
	pt2 := part2(string(input2))
	fmt.Println("The Second part result is:", pt2)
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
	return result
}
