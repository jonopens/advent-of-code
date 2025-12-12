package day_two

import (
	"fmt"
	"strconv"
	"strings"
)

type StrRange struct {
	left        string
	right       string
	IsInvalidId Invalidator
}

type Invalidator func(string) bool

func CreateStrRange(stringRep string) *StrRange {
	ends := strings.Split(stringRep, "-")
	return &StrRange{
		left:  ends[0],
		right: ends[1],
	}
}

func (sr *StrRange) GetElements() []string {
	// turn into ints so we can populate a slice with the entire range
	low, err := strconv.Atoi(sr.left)
	if err != nil {
		fmt.Println("failed to parse string to int: ", err)
	}

	high, err := strconv.Atoi(sr.right)
	if err != nil {
		fmt.Println("failed to parse string to int: ", err)
	}

	fullRange := make([]string, high-low+1)

	idx := 0
	for c := low; c < high+1; c++ {
		fullRange[idx] = strconv.Itoa(c)
		idx += 1
	}

	return fullRange
}
