package main

import (
	"fmt"
)

func fibonacci(n int) []int {
	var arr []int

	for i := 0; i < n; i++ {
		if i < 2 {
			arr = append(arr, i)
		} else {
			x := arr[i-1] + arr[i-2]
			arr = append(arr, x)
		}
	}
	return arr
}

func main() {
	var n int
	fmt.Println("enter the fibonacci series req.")
	fmt.Scanln(&n)
	fmt.Println(fibonacci(n))

}
