# Day 1

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

## Solution
1132

# Day 2

## Inputs
- probably the same
- if the inputs are the same, the solution must be greater than for part 1

## Flow
- the same with the added complexity of adding to the total any time the counter cross 0 _OR_ lands on 0