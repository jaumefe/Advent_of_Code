package main

import (
	"embed"
	"fmt"
	"log"
	"sort"
	"strconv"
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
	input2, err := inputFS.ReadFile("inputs/input.txt")
	if err != nil {
		log.Panic(err)
	}
	pt2 := part2(string(input2))
	fmt.Println("The Second part result is:", pt2)
}

type camelCard struct {
	hands []string
	bet   []int
	types []int
	rank  []int
}

func part1(input string) int {
	var result int
	var cards camelCard
	in := strings.Split(input, "\n")
	for _, i := range in {
		if len(i) != 0 {
			hand := strings.Split(i, " ")[0]
			bet, _ := strconv.Atoi(strings.Split(i, " ")[1])
			cards.bet = append(cards.bet, bet)
			cards.hands = append(cards.hands, hand)
			cards.types = append(cards.types, 0)
			cards.rank = append(cards.rank, 0)
		}
	}
	for index, i := range cards.hands {
		handType := make(map[string]int)
		for _, j := range i {
			handType[string(j)] += 1
		}
		cards.checkHandType(handType, index)
	}
	cards.sortByRank()
	cards.computeRank()
	for i := 0; i < len(cards.bet); i++ {
		result += cards.bet[i] * cards.rank[i]
	}
	return result
}

func (cards *camelCard) checkHandType(hand map[string]int, index int) {
	var dist []int
	for _, i := range hand {
		dist = append(dist, i)
	}
	sort.Ints(dist)
	var handDist string
	for _, i := range dist {
		stringTemp := strconv.Itoa(i)
		handDist += stringTemp
	}
	switch handDist {
	case "11111":
		cards.types[index] = 1
	case "1112":
		cards.types[index] = 2
	case "122":
		cards.types[index] = 3
	case "113":
		cards.types[index] = 4
	case "23":
		cards.types[index] = 5
	case "14":
		cards.types[index] = 6
	case "5":
		cards.types[index] = 7
	}
}

func (cards *camelCard) sortByRank() {
	var max, maxIndex int
	var newIndex []int
	for i := 0; i < len(cards.types); i++ {
		if cards.types[i] > max {
			max = cards.types[i]
			maxIndex = i
		}
	}
	newIndex = append(newIndex, maxIndex)
	for i := 0; i < max; i++ {
		for j := 0; j < len(cards.types); j++ {
			if j != maxIndex {
				if cards.types[j] == max-i {
					newIndex = append(newIndex, j)
				}
			}
		}
	}
	var handsTemp []string
	var betTemp, typesTemp, rankTemp []int
	for _, i := range newIndex {
		handsTemp = append(handsTemp, cards.hands[i])
		betTemp = append(betTemp, cards.bet[i])
		typesTemp = append(typesTemp, cards.types[i])
		rankTemp = append(rankTemp, cards.rank[i])
	}
	cards.hands, cards.bet, cards.types, cards.rank = handsTemp, betTemp, typesTemp, rankTemp
}

func (cards *camelCard) computeRank() {
	var rank int
	for i := 1; i <= 7; i++ {
		var sameTypesInd []int
		for j := 0; j < len(cards.types); j++ {
			if cards.types[j] == i {
				sameTypesInd = append(sameTypesInd, j)
			}
		}
		if len(sameTypesInd) != 0 {
			if len(sameTypesInd) == 1 {
				rank++
				cards.rank[sameTypesInd[0]] = rank
			} else {
				cards.highestCardCriteria(sameTypesInd, &rank, 0)
			}
		}
	}
}

func (cards *camelCard) highestCardCriteria(index []int, rank *int, start int) {
	values := map[string]int{"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "T": 10, "J": 11, "Q": 12, "K": 13, "A": 14}
	var subset []string
	for _, i := range index {
		subset = append(subset, cards.hands[i])
	}
	var subsetVal []int
	for _, i := range subset {
		subsetVal = append(subsetVal, values[string(i[start])])
	}
	maxIndex, newIndex := max(subsetVal)
	if len(maxIndex) == 1 {
		if len(newIndex) == 1 {
			*rank++
			cards.rank[index[newIndex[0]]] = *rank
			*rank++
			cards.rank[index[maxIndex[0]]] = *rank
		} else {
			var indexTemp []int
			for _, i := range newIndex {
				indexTemp = append(indexTemp, index[i])
			}
			cards.highestCardCriteria(indexTemp, rank, start)
			*rank++
			cards.rank[index[maxIndex[0]]] = *rank
		}
	} else {
		if len(newIndex) == 1 {
			*rank++
			cards.rank[index[newIndex[0]]] = *rank
			var indexTemp []int
			for _, i := range maxIndex {
				indexTemp = append(indexTemp, index[i])
			}
			cards.highestCardCriteria(indexTemp, rank, start+1)
		} else {
			var indexTemp []int
			if len(newIndex) != 0 {
				for _, i := range newIndex {
					indexTemp = append(indexTemp, index[i])
				}
				cards.highestCardCriteria(indexTemp, rank, start)
				indexTemp = make([]int, 0)
			}
			for _, i := range maxIndex {
				indexTemp = append(indexTemp, index[i])
			}
			cards.highestCardCriteria(indexTemp, rank, start+1)
		}
	}
}

func max(input []int) ([]int, []int) {
	var max int
	var maxIndex, otherIndex []int
	for _, i := range input {
		if i > max {
			max = i
		}
	}
	for index, i := range input {
		if i == max {
			maxIndex = append(maxIndex, index)
		} else {
			otherIndex = append(otherIndex, index)
		}
	}
	return maxIndex, otherIndex
}

func part2(input string) int {
	var result int
	var cards camelCard
	in := strings.Split(input, "\n")
	for _, i := range in {
		if len(i) != 0 {
			hand := strings.Split(i, " ")[0]
			bet, _ := strconv.Atoi(strings.Split(i, " ")[1])
			cards.bet = append(cards.bet, bet)
			cards.hands = append(cards.hands, hand)
			cards.types = append(cards.types, 0)
			cards.rank = append(cards.rank, 0)
		}
	}
	for index, i := range cards.hands {
		handType := make(map[string]int)
		for _, j := range i {
			handType[string(j)] += 1
		}
		cards.checkHandTypePt2(handType, index)
	}
	cards.sortByRank()
	cards.computeRankPt2()
	for i := 0; i < len(cards.bet); i++ {
		result += cards.bet[i] * cards.rank[i]
	}
	return result
}

func (cards *camelCard) checkHandTypePt2(hand map[string]int, index int) {
	var dist []int
	for _, i := range hand {
		dist = append(dist, i)
	}
	sort.Ints(dist)
	var handDist string
	for _, i := range dist {
		stringTemp := strconv.Itoa(i)
		handDist += stringTemp
	}
	switch handDist {
	case "11111":
		if hand["J"] == 1 {
			cards.types[index] = 2
			return
		}
		cards.types[index] = 1
	case "1112":
		if hand["J"] != 0 {
			cards.types[index] = 4
			return
		}
		cards.types[index] = 2
	case "122":
		if hand["J"] == 1 {
			cards.types[index] = 5
			return
		} else if hand["J"] == 2 {
			cards.types[index] = 6
			return
		}
		cards.types[index] = 3
	case "113":
		if hand["J"] != 0 {
			cards.types[index] = 6
			return
		}
		cards.types[index] = 4
	case "23":
		if hand["J"] != 0 {
			cards.types[index] = 7
			return
		}
		cards.types[index] = 5
	case "14":
		if hand["J"] != 0 {
			cards.types[index] = 7
			return
		}
		cards.types[index] = 6
	case "5":
		cards.types[index] = 7
	}
}

func (cards *camelCard) computeRankPt2() {
	var rank int
	for i := 1; i <= 7; i++ {
		var sameTypesInd []int
		for j := 0; j < len(cards.types); j++ {
			if cards.types[j] == i {
				sameTypesInd = append(sameTypesInd, j)
			}
		}
		if len(sameTypesInd) != 0 {
			if len(sameTypesInd) == 1 {
				rank++
				cards.rank[sameTypesInd[0]] = rank
			} else {
				cards.highestCardCriteriaPt2(sameTypesInd, &rank, 0)
			}
		}
	}
}

func (cards *camelCard) highestCardCriteriaPt2(index []int, rank *int, start int) {
	values := map[string]int{"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "T": 10, "J": 1, "Q": 11, "K": 12, "A": 13}
	var subset []string
	for _, i := range index {
		subset = append(subset, cards.hands[i])
	}
	var subsetVal []int
	for _, i := range subset {
		subsetVal = append(subsetVal, values[string(i[start])])
	}
	maxIndex, newIndex := max(subsetVal)
	if len(maxIndex) == 1 {
		if len(newIndex) == 1 {
			*rank++
			cards.rank[index[newIndex[0]]] = *rank
			*rank++
			cards.rank[index[maxIndex[0]]] = *rank
		} else {
			var indexTemp []int
			for _, i := range newIndex {
				indexTemp = append(indexTemp, index[i])
			}
			cards.highestCardCriteriaPt2(indexTemp, rank, start)
			*rank++
			cards.rank[index[maxIndex[0]]] = *rank
		}
	} else {
		if len(newIndex) == 1 {
			*rank++
			cards.rank[index[newIndex[0]]] = *rank
			var indexTemp []int
			for _, i := range maxIndex {
				indexTemp = append(indexTemp, index[i])
			}
			cards.highestCardCriteriaPt2(indexTemp, rank, start+1)
		} else {
			var indexTemp []int
			if len(newIndex) != 0 {
				for _, i := range newIndex {
					indexTemp = append(indexTemp, index[i])
				}
				cards.highestCardCriteriaPt2(indexTemp, rank, start)
				indexTemp = make([]int, 0)
			}
			for _, i := range maxIndex {
				indexTemp = append(indexTemp, index[i])
			}
			cards.highestCardCriteriaPt2(indexTemp, rank, start+1)
		}
	}
}
