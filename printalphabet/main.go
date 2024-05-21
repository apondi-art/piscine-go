package main

import "github.com/01-edu/z01"

func main() {
	var str string
	for i := 'a'; i <= 'z'; i++ {
		str += string(i)

	}
	var value string
	for i := len(str) - 1; i >= 0; i-- {
		if i%2 == 0 {
			value += string(str[i] - 32)

		} else {
			value += string(str[i])
		}
	}
	for _, char := range value {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')

}
