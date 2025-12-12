package day_two

import (
	"advent_of_code/internal/constants"
	"advent_of_code/internal/utils"
	"fmt"
	"strconv"
)

func areHalvesEqual(id string) bool {
	middle := len(id) / 2
	left, right := id[:middle], id[middle:]

	if len(left) != len(right) {
		return false
	}
	return left == right
}

func DayTwoPartOne() int {
	total := 0

	ranges, err := utils.GetLinesFromSingleLineBySeparator(constants.DayTwo2025InputPath, ",")
	if err != nil {
		fmt.Println("failed to parse inputs from file: ", err)
	}

	for _, r := range ranges {
		// make into a list of elements from the string representation of the range
		// they should be strings for ease of comparison of left and right
		strRange := CreateStrRange(r)
		strRange.IsInvalidId = areHalvesEqual

		fullRange := strRange.GetElements()
		for _, el := range fullRange {
			isInvalid := strRange.IsInvalidId(el)
			if isInvalid && el != "" {
				toAdd, newErr := strconv.Atoi(el)
				if newErr != nil {
					fmt.Println("failed to convert to int: ", newErr, toAdd, el)
				}

				total += toAdd
			}

		}
	}
	fmt.Println("d2p1 solution: ", total)

	return total
}
