package main

import "fmt"

func main() {
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)

}
func Chunk(t []int, size int) {
	var chunk [][]int
	n := size

	for i := 0; i < len(t); i += size { //value of i increases with  i+=size
		if n < len(t) {
			chunk = append(chunk, t[i:n])
		} else {
			chunk = append(chunk, t[i:])
		}
		n += size

	}
	fmt.Println(chunk)
}
