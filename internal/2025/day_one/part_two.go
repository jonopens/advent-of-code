package day_one

import (
	"advent_of_code/internal/constants"
	"advent_of_code/internal/utils"
	"fmt"
)

func DayOnePartTwo() int {
	score := 0
	safeDial := CreateSafeDial(50, 99)

	inputs, err := utils.GetLinesFromInput(constants.DayOne2025InputPath)
	if err != nil {
		fmt.Println("Failed to parse file: ", err)
		return 0
	}

	for idx, input := range inputs {
		move, err := CreateMove(input)
		if err != nil {
			fmt.Println("failed to create move! --> ", err)
		}

		oldPosition := safeDial.Position
		score += move.FullTurns() // do this before changing position
		signedClicks := move.ClicksSigned()

		safeDial.Turn(signedClicks)
		if signedClicks > 0 && oldPosition+(signedClicks%100) > 100 {
			score += 1
		}

		if signedClicks < 0 && safeDial.Position > oldPosition {
			score += 1
		}

		if safeDial.Position%100 == 0 {
			score += 1
		}

		if idx < 10 {
			fmt.Println("summary:")
			fmt.Println("  move-distance: ", move.Clicks, " move-signed: ", move.ClicksSigned())
			fmt.Println("  old position: ", oldPosition)
			fmt.Println("  new position: ", safeDial.Position)
			fmt.Println("  score: ", score)
		}
	}

	fmt.Println("solution: ", score)
	return score
}
