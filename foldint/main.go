package main

import "fmt"

func FoldInt(f func(int, int) int, a []int, n int) {
	table := 93
	for i := 0; i <= len(a)-1; i++ {
		table = f(table, a[i])
	}
	fmt.Println(table)

}
func main() {

	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}
func Add(a, b int) int {
	return a + b
}
func Mul(a, b int) int {
	return a * b
}
func Sub(a, b int) int {
	return a - b
}
