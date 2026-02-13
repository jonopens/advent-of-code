# Day 3, Part 1

## Inputs
- lines of digits

## Flow
- for each line in the file, use two numeric characters to form the two digit number, in order
-- iterate through n-2 (tens place)
-- find greatest and store the value and its index
-- starting from idx +1 until n-1, iterate to find the greatest ones value
-- NOTE: all > must be using byte values because iterating on a string returns the values for the runes

# Day 2, Part 2

## Adjusted Flow
- Nested for called on each row
- function adjusts based on target number of batteries, in this case 12

## Self-review: Could It Be Faster?
- it's cubic so that feels bad
- first for loop cannot be eliminated, so could it be quadratic?
- I cannot conceive of a linear solution
