package main

import (
	"github.com/01-edu/z01"
	//"github.com/01-edu"
)

func IsNegative(nb int) {
	if nb >= 0 {
		z01.PrintRune('F')
	} else {
		z01.PrintRune('T')
	}
	z01.PrintRune('\n')

}
func main() {
	IsNegative(-34)
	IsNegative(0)
	IsNegative(-1)
}
