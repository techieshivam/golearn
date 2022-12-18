package main

import (
	"fmt"
)

func main() {
	var n, r int
	fmt.Println("enter the no.")
	fmt.Scanln(&n)
	sum := 0
	for n > 0 {
		r = n % 10
		sum = sum + r
		n = n / 10

	}
	fmt.Println("sum of digits of no is", sum)
}
