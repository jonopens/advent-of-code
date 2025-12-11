package day_one

import (
	"advent_of_code/internal/constants"
	"advent_of_code/internal/utils"
	"fmt"
)

func DayOnePartTwo() int {
	score := 0
	safeDial := CreateSafeDial(50, 99)

	inputs, err := utils.GetLinesFromMultilineInput(constants.DayOne2025InputPath)
	if err != nil {
		fmt.Println("Failed to parse file: ", err)
		return 0
	}

	for _, input := range inputs {
		zeroCrossings := 0
		move, err := CreateMove(input)
		if err != nil {
			fmt.Println("failed to create move! --> ", err)
		}

		oldPosition := safeDial.Position
		fullTurns := move.FullTurns()
		zeroCrossings += fullTurns
		signedClicks := move.ClicksSigned()

		safeDial.Turn(signedClicks)
		if signedClicks > 0 && oldPosition+(signedClicks%100) > 100 {
			zeroCrossings += 1
		}

		if signedClicks < 0 && (safeDial.Position > oldPosition && oldPosition != 0) {
			zeroCrossings += 1
		}

		if safeDial.Position%100 == 0 {
			zeroCrossings += 1
		}

		// make first arg idx if debugging
		// if (idx > 100 && idx < 200) && (move.ClicksSigned() > 100 || move.ClicksSigned() < -100) {
		// 	fmt.Println("summary @: ", idx)
		// 	fmt.Println("  move-signed: ", move.ClicksSigned())
		// 	fmt.Println("  full turns: ", fullTurns)
		// 	fmt.Println("  old position: ", oldPosition)
		// 	fmt.Println("  new position: ", safeDial.Position)
		// 	fmt.Println("  zero crossing: ", zeroCrossings)
		// 	fmt.Println("  score: ", score)
		// }

		// if oldPosition == 0 {
		// 	fmt.Println("summary @: ", idx)
		// 	fmt.Println("  move-signed: ", move.ClicksSigned())
		// 	fmt.Println("  full turns: ", fullTurns)
		// 	fmt.Println("  old position: ", oldPosition)
		// 	fmt.Println("  new position: ", safeDial.Position)
		// 	fmt.Println("  zero crossing: ", zeroCrossings)
		// 	fmt.Println("  score: ", score)
		// }

		score += zeroCrossings
	}

	fmt.Println("d1p2 solution: ", score)
	return score
}
