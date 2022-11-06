package main

import (
	"fmt"
)

func main() {
	var n1, n2 int
	fmt.Println("enter your first no.")
	fmt.Scanln(&n1)
	fmt.Println("enter your second no.")
	fmt.Scanln(&n2)

	sum := 0
	for i := n1 + 1; i < n2; i++ {
		sum = sum + i
	}
	fmt.Println("sum of number between no. is", sum)
}
