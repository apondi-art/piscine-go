package main

import (
	"fmt"
	"os"
)

func AphaMirror(s string) string {
	var str string
	var mirror rune
	var charindex rune
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			//(char-a) getting the index position of the character in aphabet stating from index 0 which is 'a'
			charindex = (char - 'a')
			mirror = 'z' - charindex // getting mirror value from posurion z

		} else if char >= 'A' && char <= 'Z' {
			charindex = char - 'A'
			mirror = 'Z' - charindex

		} else {
			mirror = char
		}
		str += string(mirror)
	}
	return str
}
func main() {
	argument := os.Args[1]
	fmt.Println(AphaMirror(argument))
}
