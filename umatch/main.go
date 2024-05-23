package main

import (
	"fmt"
	"os"
)

func main() {
	
	str1 := os.Args[1]
	str2 := os.Args[2]
	i := 0
	for _, char := range str2 {
		if char == rune(str1[i]) {
			i++
		}
		if i == len(str1) {
			fmt.Println(str1)
			return
		}
	}
}
