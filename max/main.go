package main

import (
	"github.com/01-edu/z01"
)

func Max(t []int) int {
	max := 0
	for i := 0; i <= len(t)-1; i++ {
		if t[i] > max {
			max = t[i]
		}
	}
	return max
}
func main() {
	a := []int{23, 123, 1, 11, 55, 93, 478}
	max := Max(a)
	n := Itoa(max)
	for _, run := range n {
		z01.PrintRune(run)

	}
	z01.PrintRune('\n')

}
func Itoa(nb int) string {
	var mod int
	var run rune
	var div int
	var str string
	div = nb
	for div > 0 {
		mod = div % 10
		div = div / 10
		run = int32(mod + '0')
		str = string(run) + str

	}
	return str
}
