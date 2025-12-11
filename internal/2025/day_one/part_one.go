package day_one

import (
	"advent_of_code/internal/constants"
	"advent_of_code/internal/utils"
	"fmt"
)

func DayOnePartOne() int {
	score := 0
	safeDial := CreateSafeDial(50, 99)

	inputs, err := utils.GetLinesFromMultilineInput(constants.DayOne2025InputPath)
	if err != nil {
		fmt.Println("Failed to parse file: ", err)
		return 0
	}

	for _, input := range inputs {
		move, err := CreateMove(input)
		if err != nil {
			fmt.Println("Failed to create move! --> ", err)
		}

		safeDial.Turn(move.ClicksSigned())
		if safeDial.Position%100 == 0 {
			score = score + 1
		}
	}

	fmt.Println("d1p1 solution: ", score)
	return score
}
