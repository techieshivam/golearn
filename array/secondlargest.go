package main

import (
	"fmt"
)

func main() {
	var size int
	fmt.Println("enter the size of array")
	fmt.Scanln(&size)

	if size == 0 {
		fmt.Println("No action can be performed!")
		return
	}

	array := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Printf("Enter value of array[%d] = ", i)
		fmt.Scanln(&array[i])
	}

	fmt.Println("Array is :", array)

	large1 := 0
	large2 := 0

	large1 = array[0]
	for i := 1; i <= size-1; i++ {
		if large1 < array[i] {
			large2 = large1
			large1 = array[i]
		} else if large2 < array[i] {
			large2 = array[i]
		}
	}
	fmt.Println("Second largest element is: ", large2)
}
