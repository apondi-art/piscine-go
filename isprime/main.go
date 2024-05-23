package main

import "fmt"

func Isprime(nb int) bool {
	var count int
	for i := 1; i <= nb; i++ {
		if nb%i == 0 {
			count++
		}
	}
	if count == 2 {
		return true
	} else {
		return false
	}
}
func main() {
	nb := 15
	fmt.Println(PreviousPrime(nb))
	fmt.Println(Isprime(15))
}
func PreviousPrime(current int)int{
	var prime int
	
	for i:= current; i>= 1; i--{
		if Isprime(i){
			prime = i
			break
		}
		
	}
	return prime
}
