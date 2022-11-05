package main

import "fmt"

func main() {

	a := []int{4, 3, 6, 8, 3, 2, 4, 4, 9}
	fmt.Println("Given array is", a)

	var result = make(map[int]int)

	for _, num := range a {

		result[num] = result[num] + 1

	}
	fmt.Println("Final result", result)
}
