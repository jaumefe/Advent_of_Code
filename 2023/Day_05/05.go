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
	input2, err := inputFS.ReadFile("inputs/test.txt")
	if err != nil {
		log.Panic(err)
	}
	pt2 := part2(string(input2))
	fmt.Println("The Second part result is:", pt2)
}

func part1(input string) int {
	var result int
	in := strings.Split(input, "\n")
	var data []string
	for _, i := range in {
		if len(i) != 0 {
			data = append(data, i)
		}
	}
	var seeds, seedToSoil, soilToFer, ferToWat, watToLight, lightToTemp, tempToHum, humToLoc []string
	var ss, sf, fw, wl, lt, th, hl int
	for index, i := range data {
		if index == 0 {
			seeds = append(seeds, i)
		} else {
			if i == "seed-to-soil map:" {
				ss = 1
			} else if i == "soil-to-fertilizer map:" {
				sf = 1
				ss = 0
			} else if i == "fertilizer-to-water map:" {
				fw = 1
				sf = 0
			} else if i == "water-to-light map:" {
				wl = 1
				fw = 0
			} else if i == "light-to-temperature map:" {
				lt = 1
				wl = 0
			} else if i == "temperature-to-humidity map:" {
				th = 1
				lt = 0
			} else if i == "humidity-to-location map:" {
				hl = 1
				th = 0
			}
		}
		if ss == 1 {
			seedToSoil = append(seedToSoil, i)
		} else if sf == 1 {
			soilToFer = append(soilToFer, i)
		} else if fw == 1 {
			ferToWat = append(ferToWat, i)
		} else if wl == 1 {
			watToLight = append(watToLight, i)
		} else if lt == 1 {
			lightToTemp = append(lightToTemp, i)
		} else if th == 1 {
			tempToHum = append(tempToHum, i)
		} else if hl == 1 {
			humToLoc = append(humToLoc, i)
		}

	}
	seedsIndex := strArrToIntArr(strings.Split(strings.Split(seeds[0], ": ")[1], " "))
	seedSoilMap, soilFerMap, ferWatMap, watLightMap, lightTempMap, tempHumMap, humLocMap := fillMapsData(seedToSoil), fillMapsData(soilToFer), fillMapsData(ferToWat), fillMapsData(watToLight), fillMapsData(lightToTemp), fillMapsData(tempToHum), fillMapsData(humToLoc)
	soil := translateFromMapToMap(seedsIndex, seedSoilMap)
	fer := translateFromMapToMap(soil, soilFerMap)
	wat := translateFromMapToMap(fer, ferWatMap)
	light := translateFromMapToMap(wat, watLightMap)
	temp := translateFromMapToMap(light, lightTempMap)
	hum := translateFromMapToMap(temp, tempHumMap)
	loc := translateFromMapToMap(hum, humLocMap)
	sort.Ints(loc)
	result = loc[0]
	return result
}

type maps struct {
	dest   int
	source int
	length int
}

func translateFromMapToMap(input []int, mapData []maps) []int {
	var result []int
	for _, i := range input {
		var count int
		for _, j := range mapData {
			if i >= j.source && i < j.source+j.length {
				result = append(result, j.dest+(i-j.source))
				break
			} else {
				count++
			}
		}
		if count == len(mapData) {
			result = append(result, i)
		}
	}
	return result
}

func fillMapsData(input []string) []maps {
	var result []maps
	var input_temp []string
	var inputInt [][]int

	for i := 1; i < len(input); i++ {
		input_temp = append(input_temp, input[i])
	}
	for _, i := range input_temp {
		var inputIntTemp []int
		var maps_temp maps
		for _, j := range strings.Split(i, " ") {
			num, err := strconv.Atoi(j)
			if err != nil {
				log.Fatal(err)
			}
			inputIntTemp = append(inputIntTemp, num)
		}
		result = append(result, maps_temp)
		inputInt = append(inputInt, inputIntTemp)
	}
	for index, i := range inputInt {
		result[index].dest = i[0]
		result[index].source = i[1]
		result[index].length = i[2]
	}
	return result
}

func part2(input string) int {
	var result int
	in := strings.Split(input, "\n")
	var data []string
	for _, i := range in {
		if len(i) != 0 {
			data = append(data, i)
		}
	}
	var seeds, seedToSoil, soilToFer, ferToWat, watToLight, lightToTemp, tempToHum, humToLoc []string
	var ss, sf, fw, wl, lt, th, hl int
	for index, i := range data {
		if index == 0 {
			seeds = append(seeds, i)
		} else {
			if i == "seed-to-soil map:" {
				ss = 1
			} else if i == "soil-to-fertilizer map:" {
				sf = 1
				ss = 0
			} else if i == "fertilizer-to-water map:" {
				fw = 1
				sf = 0
			} else if i == "water-to-light map:" {
				wl = 1
				fw = 0
			} else if i == "light-to-temperature map:" {
				lt = 1
				wl = 0
			} else if i == "temperature-to-humidity map:" {
				th = 1
				lt = 0
			} else if i == "humidity-to-location map:" {
				hl = 1
				th = 0
			}
		}
		if ss == 1 {
			seedToSoil = append(seedToSoil, i)
		} else if sf == 1 {
			soilToFer = append(soilToFer, i)
		} else if fw == 1 {
			ferToWat = append(ferToWat, i)
		} else if wl == 1 {
			watToLight = append(watToLight, i)
		} else if lt == 1 {
			lightToTemp = append(lightToTemp, i)
		} else if th == 1 {
			tempToHum = append(tempToHum, i)
		} else if hl == 1 {
			humToLoc = append(humToLoc, i)
		}

	}
	seedsIndex := strArrToIntArr(strings.Split(strings.Split(seeds[0], ": ")[1], " "))
	var newSeeds []int

	for i := 0; i < len(seedsIndex); i += 2 {
		for j := seedsIndex[i]; j < seedsIndex[i]+seedsIndex[i+1]; j++ {
			newSeeds = append(newSeeds, j)
		}
	}
	//fmt.Println(newSeeds)
	seedSoilMap, soilFerMap, ferWatMap, watLightMap, lightTempMap, tempHumMap, humLocMap := fillMapsData(seedToSoil), fillMapsData(soilToFer), fillMapsData(ferToWat), fillMapsData(watToLight), fillMapsData(lightToTemp), fillMapsData(tempToHum), fillMapsData(humToLoc)
	soil := translateFromMapToMap(newSeeds, seedSoilMap)
	fer := translateFromMapToMap(soil, soilFerMap)
	wat := translateFromMapToMap(fer, ferWatMap)
	light := translateFromMapToMap(wat, watLightMap)
	temp := translateFromMapToMap(light, lightTempMap)
	hum := translateFromMapToMap(temp, tempHumMap)
	loc := translateFromMapToMap(hum, humLocMap)
	sort.Ints(loc)
	result = loc[0]
	return result
}

func strArrToIntArr(input []string) []int {
	res := make([]int, len(input))
	for i, j := range input {
		res[i], _ = strconv.Atoi(j)
	}
	return res
}
