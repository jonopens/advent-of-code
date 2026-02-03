# Day 1, Part 1

## Inputs
- each input line provides:
-- a direction/operator (L => subtraction, R => addition)
-- an number/operand

## Flow
- parse the input text file into a slice of strings
- iterate and break string into operator and quantity
- add or subtract
- increment a counter any time the ABS of the total modulus 100 is 0
- return the counter

# Day 1, Part 2

## Adjusted Flow
- the same with the added complexity of incrementing the total any time the counter cross 0 _OR_ lands on 0

## Self-review: Could It Be Faster?
- TBD
