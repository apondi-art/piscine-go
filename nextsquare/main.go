package main

import (
	"fmt"
	"math"
)

func main() {
	y := 121
	s := NextSquare(int64(y))
	fmt.Println(s)
}
func NextSquare(number int64) int64 {
	if number < 0 { // checks if number is negative
		return 0
	}
	squareroot := math.Sqrt(float64(number)) //getting squareroot of number
	if math.Mod(squareroot, 1) != 0 {        // checking if the squareroot  obtained is a whole number using math.Mod(squareroot ,1)
		return -1
	}
	return int64(squareroot+1) * int64(squareroot+1) // return squareroot + 1 * squareroot + 1
}
