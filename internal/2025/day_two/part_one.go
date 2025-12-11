package day_two

import (
	"advent_of_code/internal/constants"
	"advent_of_code/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

func getRangeSlice(start string, end string) []string {
	low, err := strconv.Atoi(start)
	if err != nil {
		fmt.Println("failed to parse string to int: ", err)
	}

	high, err := strconv.Atoi(end)
	if err != nil {
		fmt.Println("failed to parse string to int: ", err)
	}

	fullRange := make([]string, high-low+1)

	for c := low; c <= high+1; c++ {
		fullRange = append(fullRange, strconv.Itoa(c))
	}

	return fullRange
}

func isInvalidId(id string) bool {
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
		ends := strings.Split(r, "-")
		fullRange := getRangeSlice(ends[0], ends[1])
		for _, el := range fullRange {
			isInvalid := isInvalidId(el)
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
