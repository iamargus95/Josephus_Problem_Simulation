# Josephus_Problem_Simulation

Solve Josephus problem using the 2l+1 method.

## Implementation

1. Get input integer from CLI
2. Find largest power of 2 inside of the above integer.
3. Substract the largest power of 2 from input integer and assign it to 'l'.
4. The Survivor is standing at the position (2*l)+1.

## USAGE:

`go run main.go [integer]`

eg. `go run main.go 91` etc.