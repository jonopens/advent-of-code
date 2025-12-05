package day_one

import (
	"fmt"
	"strconv"
)

type Move struct {
	direction string
	Clicks    int
}

func CreateMove(moveText string) (*Move, error) {
	dir, count := moveText[:1], moveText[1:]
	countInt, err := strconv.Atoi(count) // int
	if err != nil {
		return &Move{}, fmt.Errorf("failed to convert string to int: %f", err)
	}

	return &Move{
		direction: dir,
		Clicks:    countInt,
	}, nil
}

func (m *Move) ClicksSigned() int {
	if m.direction == "L" {
		return -m.Clicks
	}

	return m.Clicks
}

func (m *Move) FullTurns() int {
	return m.Clicks / 100
}
