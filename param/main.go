package main

import (
	"os"
)

func main() {
	if len(os.Args) < 2 {
		os.Stdout.WriteString("legth of argument is less" + "\n")
		return
	}
	argument := os.Args[1:]

	os.Stdout.WriteString(argument[len(argument)-1] + "\n")

}
