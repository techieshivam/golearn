package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("enter the number")
	fmt.Scanln(&n)
	fmt.Println("n is", n)
	n1 := n
	var r int
	rev := 0
	for n > 0 {
		r = n % 10
		rev = rev*10 + r
		n = n / 10
	}
	fmt.Println("rev", rev)

	if rev == n1 {
		fmt.Println("It is Pellindrome number")
	} else {
		fmt.Println("Not a Pellindrome number")
	}

}
