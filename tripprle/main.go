package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abCd"
	val := Tripple(s)
	fmt.Println(val)
}

func Tripple(s string) string {
	var str string
	// var value string
	for i, char := range s {
		for j := 0; j < i+1; j++ {
			if j == 0 {
				str += strings.ToUpper(string(char))
			} else if char >= 'A' && char <= 'Z' {
				str += strings.ToLower(string(char))
			} else {
				str += string(char)
			}
		}
		if i != len(s)-1 {
			str += "-"
		}

	}
	return str
}
