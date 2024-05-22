package main

import (
	"fmt"
	"os"
)

//function that returns length of arguments passed on terminal"
func main() {
	arguments := os.Args[1:]
	args := len(arguments)
	fmt.Println(args)
}
