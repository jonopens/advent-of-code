package day_one

import (
	"advent_of_code/constants"
	"advent_of_code/utils"
	"fmt"
	"strconv"
)

func GetMove(moveDesc string) int {
	dir, count := moveDesc[:1], moveDesc[1:]
	countInt, convErr := strconv.Atoi(count)
	if convErr != nil {
		return 0
	}

	fmt.Println("current move: ", dir, count)
	if dir == "L" {
		return -countInt
	}
	return countInt
}

func DayOnePartOne() int {
	score := 0
	position := 50

	inputs, err := utils.GetLinesFromInput(constants.DayOne2025InputPath)
	if err != nil {
		fmt.Println("Failed to parse file: ", err)
		return 0
	}

	for _, input := range inputs {
		position += GetMove(input)
		fmt.Println("current position: ", position)
		if utils.AbsInt(position)%100 == 0 {
			fmt.Println("---")
			fmt.Println("increasing score: ", score)
			score = score + 1
		}
	}

	fmt.Println("solution: ", score)
	return score
}
