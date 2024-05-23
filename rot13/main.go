package main

import (
	"fmt"
	"os"
)

func Rot13(s string) string {
	var run rune
	var str string
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			// Convert char to 0-25 index, apply ROT13, then modulo 26 to wrap around
			run = 'a' + (char-'a'+13)%26
		} else if char >= 'A' && char <= 'Z' {
			// Convert char to 0-25 index, apply ROT13, then modulo 26 to wrap around
			run = 'A' + (char-'A'+13)%26
		} else {
			// Non-alphabetic characters remain unchanged
			run = char
		}
		str += string(run)
	}
	return str
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a string to encode.")
		return
	}
	str := os.Args[1]
	fmt.Println(Rot13(str))
}
