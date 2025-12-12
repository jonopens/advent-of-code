package day_two

import (
	"advent_of_code/internal/constants"
	"advent_of_code/internal/utils"
	"fmt"
	"strconv"
)

func allValuesEqual[T comparable](s []T) bool {
	if len(s) == 0 {
		fmt.Println("zero length slice condition!")
		return true
	}

	first := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] != first {
			return false
		}
	}

	return true
}

func repeatsEqualLengthChunks(id string) bool {
	containsRepetition := false
	length := len(id)

	for i := 1; i <= (length / 2); i++ { //
		if length%i != 0 {
			continue // skip loops where full repetition isn't possible
			// because the string cannot be equally divided by i
		}

		chunks := make([]string, length/i)
		for j := 0; j < length/i; j++ {
			start := i * j
			end := (i * j) + i
			chunks[j] = id[start:end]
			// fmt.Println("chunks in repeatsEqualLengthChunks: ", chunks)
		}

		eqCheck := allValuesEqual(chunks)
		if eqCheck {
			containsRepetition = allValuesEqual(chunks)
			break
		}
	}
	return containsRepetition
}

func DayTwoPartTwo() int {
	total := 0

	ranges, err := utils.GetLinesFromSingleLineBySeparator(constants.DayTwo2025InputPath, ",")
	if err != nil {
		fmt.Println("failed to parse inputs from file: ", err)
	}

	for _, r := range ranges {
		// make into a list of elements from the string representation of the range
		// they should be strings for ease of comparison of left and right
		strRange := CreateStrRange(r)
		strRange.IsInvalidId = repeatsEqualLengthChunks

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
	fmt.Println("d2p2 solution: ", total)

	return total
}
