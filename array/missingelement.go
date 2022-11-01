package main

import (
	"fmt"
)

func main() {
	var array [6]int
	array[0] = 1
	array[1] = 2
	array[2] = 4
	array[3] = 5
	array[4] = 6
	array[5] = 7
	fmt.Println(array)
	fmt.Println(len(array))
	a := len(array) + 1
	b := len(array) + 2
	sum := (a * b) / 2
	fmt.Println(sum)
	sum1 := 0
	for i := 0; i < len(array); i++ {
		sum1 = sum1 + array[i]
	}
	fmt.Println("sum is", sum1)

	fmt.Println("Missing no. in Array is", sum-sum1)

}
