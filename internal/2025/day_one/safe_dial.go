package day_one

type SafeDial struct {
	Position int
	Max      int
}

func CreateSafeDial(startingPosition int, max int) *SafeDial {
	return &SafeDial{
		Position: startingPosition,
		Max:      max,
	}
}

func (s *SafeDial) circularIndex(position int) int {
	position = position % (s.Max + 1) // assign the modulus of the total # of positions e.g 100
	if position < 0 {
		// when the new position is negative, invert by adding total positions
		// e.g. -15 -> 85, aka 15 back from 0, which is correct
		position += (s.Max + 1)
	}

	return position
}

func (s *SafeDial) Turn(clicks int) {
	s.Position = s.circularIndex(clicks + s.Position)
}
