package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := []any{"5", "0", 9, 3, 2, 1, "9", 6, 7}
	v := SumMix(s)
	fmt.Println(v)
}
//func takes any type and return sum of the value passed
func SumMix(arr []any) (sum int) {
	for _, element := range arr {
		if str, ok := element.(string); ok {
			el, _ := strconv.Atoi(str)
			sum += el
		} else if num, ok := element.(int); ok {
			sum += num
		}
	}
	return
}
