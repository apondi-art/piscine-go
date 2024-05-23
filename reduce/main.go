package main

func ReduceInt(a []int, f func(int, int) int) int {
	res := a[0]
	for i := 1; i < len(a); i++ {
		res = f(res, a[i])

	}
	return res
}
func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2, 45, 7, 9}
	println(ReduceInt(as, mul))
	println(ReduceInt(as, sum))
	println(ReduceInt(as, div))
}
