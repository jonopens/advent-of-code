package day_three

import (
	"advent_of_code/internal/constants"
	"advent_of_code/internal/utils"
	"fmt"
	"strconv"
)

func getMaxJoltsForNBatteries(row string, num int) string {
	jolt := ""
	lastJoltIndex := 0
	interiorLoopIndex := 0
	compareByte := byte('0')
	for i := num; i >= 1; i-- {
		// each i has it's own for iteration
		// the lower limit is the last jolt index
		// the constraint is the row length minute i
		// if it's greater, set compare byte and update jolt index
		// after that, if it's the last index
		// add to the jolt string and reset compare byte
		if lastJoltIndex != 0 {
			interiorLoopIndex = lastJoltIndex + 1
		}

		for j := interiorLoopIndex; j <= len(row)-i; j++ {
			if row[j] > compareByte {
				compareByte = row[j]
				lastJoltIndex = j
			}

			if j == len(row)-i {
				jolt += string(compareByte)
				compareByte = byte('0')
			}
		}
	}

	return jolt
}

// func test(row string, batteryCount int) string {
// 	tens := byte('0') // initialize as zero bytes for comparison
// 	ones := byte('0')
// 	tensIndex := 0

// 	for i := 0; i <= len(row)-2; i++ { // iterate only for tens position
// 		if row[i] > tens {
// 			tens = row[i]
// 			tensIndex = i
// 		}
// 	}

// 	for j := tensIndex + 1; j <= len(row)-1; j++ {
// 		if row[j] > ones {
// 			ones = row[j]
// 		}
// 	}

// 	lineMax := string(tens) + string(ones)

// 	return lineMax
// }

func DayThreePartTwo() int {
	total := 0

	lines, err := utils.GetLinesFromMultilineInput(constants.DayThree2025InputPath)
	if err != nil {
		fmt.Println("failed to get lines from file: ", err)
	}

	for idx, bank := range lines {
		fmt.Println("line no.: ", idx+1)
		fmt.Println("full bank: ", bank)
		lineMax := getMaxJoltsForNBatteries(bank, 12)

		fmt.Println("lineMax: ", lineMax)
		fmt.Println("")

		lineMaxInt, interr := strconv.Atoi(lineMax)
		fmt.Println("total will be ", total, "+ ", lineMax, "which is ", total+lineMaxInt)
		fmt.Println("")
		if err != nil {
			fmt.Println("failed to convert string to int: ", interr)
		}
		total += lineMaxInt
	}

	fmt.Println("d3p1 solution: ", total)

	return total
}
