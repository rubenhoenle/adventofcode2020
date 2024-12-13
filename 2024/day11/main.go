package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rubenhoenle/adventofcode/2024/utils"
)

func parseStoneInput(input []string) ([]int, error) {
	stones := []int{}
	for _, line := range input {
		for _, subStr := range strings.Split(line, " ") {
			stoneVal, err := strconv.Atoi(subStr)
			if err != nil {
				return []int{}, err
			}
			stones = append(stones, stoneVal)
		}
	}
	return stones, nil
}

func getStonesMap(stones []int) map[int]int {
	stoneMap := map[int]int{}
	for _, stone := range stones {
		// if key exists in the map, increment the value
		if _, ok := stoneMap[stone]; ok {
			stoneMap[stone] += 1
			continue
		}
		// if key does not exist in the map, initialize it
		stoneMap[stone] = 1
	}
	return stoneMap
}

// returns the stone map after N blinks
func blinkNTimes(n int, stoneMap map[int]int) map[int]int {
	currentMap := make(map[int]int)
	for k, v := range stoneMap {
		currentMap[k] = v
	}

	for i := 0; i < n; i++ {
		// we need a second temporary map, to not mess up the map we're iterating over
		nextMap := map[int]int{}

		for stoneVal, stoneCount := range currentMap {
			if stoneVal == 0 {
				// TODO maybe we can find a way to get rid of the duplicated code here
				if _, ok := nextMap[1]; ok {
					nextMap[1] += stoneCount
				} else {
					nextMap[1] = stoneCount
				}
			} else if len(strconv.Itoa(stoneVal))%2 == 0 {
				stoneValStr := strconv.Itoa(stoneVal)
				leftHalf, _ := strconv.Atoi(stoneValStr[:len(stoneValStr)/2])
				rightHalf, _ := strconv.Atoi(stoneValStr[len(stoneValStr)/2:])
				for _, half := range []int{leftHalf, rightHalf} {
					if _, ok := nextMap[half]; ok {
						nextMap[half] += stoneCount
					} else {
						nextMap[half] = stoneCount
					}
				}
			} else {
				newVal := stoneVal * 2024
				if _, ok := nextMap[newVal]; ok {
					nextMap[newVal] += stoneCount
				} else {
					nextMap[newVal] = stoneCount
				}
			}
		}

		// update the map for the next run
		currentMap = make(map[int]int)
		for k, v := range nextMap {
			currentMap[k] = v
		}
	}
	return currentMap
}

func main() {
	input := utils.ReadLinesFromFile("input/11.txt")
	stones, _ := parseStoneInput(input)
	stoneMap := getStonesMap(stones)

	part1 := utils.SumMapValues(blinkNTimes(25, stoneMap))
	fmt.Printf("Part 1: %d\n", part1)

	part2 := utils.SumMapValues(blinkNTimes(75, stoneMap))
	fmt.Printf("Part 2: %d\n", part2)
}
