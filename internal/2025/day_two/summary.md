# Day 2, Part 1

## Inputs
- single line of comma-separated values

## Flow
- parse input into ranges
- for each number in range
-- split in the middle and compare halves for equality
-- if equal, add element to list of matches and add value to solution total

# Day 2, Part 2

## Inputs
- the same

## Flow
- instead of checking equal halves, we are checking for equally divisible repition
-- ints with prime length are automatically valid (not added)
-- if the total len isn't cleanly divisible by the substr len, it's valid (not added)
-- single digit repetitions are invalid (added)
-- must be equally divisible