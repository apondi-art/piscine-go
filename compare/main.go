package main

import (
	"fmt"
)

func Compare(str1, str2 string) int {
	if str1 == str2 {
		return 0
	}
	if str1 < str2 {
		return -1
	}
	return 1

}
func main() {

	fmt.Println(Compare("Hello!", "Hello!"))

	fmt.Println(Compare("Ola!", "Ol"))
	fmt.Println(Compare("z", "B"))

}
