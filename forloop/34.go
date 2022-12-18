package main

import (
	"fmt"
)

func main() {
	var n1, n2, d int

	fmt.Println("enter the first no.")
	fmt.Scanln(&n1)
	fmt.Println("enter the second no.")
	fmt.Scanln(&n2)
	fmt.Println("enter the divisor")
	fmt.Scanln(&d)

	sum := 0
	for i := n1 + 1; i < n2; i++ {
		if i%d == 0 {
			sum = sum + i
		}

	}
	fmt.Println("required answer is", sum)

}
