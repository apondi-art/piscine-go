package main

import "fmt"

func main() {
	s := "abcdef"
	if len(s)%2 == 0 {
		//checks if len ofstring is even
		midpoint := len(s) / 2
		fmt.Println(string(s[midpoint-1]) + string(s[midpoint]))

	} else {
		//length of string given is odd
		midpoint := (len(s) - 1) / 2
		fmt.Println(string(s[midpoint]))
	}
}
