package main

import "fmt"

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

	var count int
	for i := 0; i <= 5; i++ {
		count = 0
		for j := 2; j < array[i]/2; j++ {
			if array[i]%j == 0 {
				count++
			}
		}
		if count == 0 && array[i] > 1 {
			fmt.Println("prime number in array ", array[i])
		}
	}
}
