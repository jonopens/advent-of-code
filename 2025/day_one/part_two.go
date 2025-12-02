package day_one

import (
	"advent_of_code/constants"
	"advent_of_code/utils"
	"fmt"
)

func GetZeroCrossings(one int, two int) int {

}

func DayOnePartTwo() int {
	score := 0
	position := 50

	inputs, err := utils.GetLinesFromInput(constants.InputPath)
	if err != nil {
		fmt.Println("Failed to parse file: ", err)
		return 0
	}

	for _, input := range inputs {
		move := day_one.GetMove(input)
	}

	fmt.Println("current score: ", score)
	return score
}
