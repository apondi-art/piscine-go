package main

import (
	"fmt"
	"os"
)

func Atoi(s string) int {
	var result int
	var n int
	for _, char := range s {
		n = int(char - '0')
		result = result*10 + n
	}
	return result
}
func main() {
	var result int
	str1 := os.Args[1]
	str2 := os.Args[2]
	str3 := os.Args[3]
	if str2 == "*" {
		result = Atoi(str1) * Atoi(str3)
	} else if str2 == "/" {
		result = Atoi(str1) / Atoi(str3)
	} else if str2 == "+" {
		result = Atoi(str1) + Atoi(str3)
	} else if str2 == "-" {
		result = Atoi(str1) - Atoi(str3)

	} else if str2 == "%" {
		result = Atoi(str1) % Atoi(str3)
	}
	fmt.Println(result)

}
