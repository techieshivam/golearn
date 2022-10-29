// 1. Program (*USE: Functions, Switch cases and it should be menu based)
// (i)  Prints array element
// (ii) Length of an array
// (iii) Display array in reverse order
// (iv) Maximum element from an array ({1,2,3,6,5} ==> 6)
// (v)	 Minimum element from an array
// (vi) Sum of all elements of an array
// (vii) Copy array elements to other array
// (viii) Print duplicate elements from an array
// (ix) Unique element from an array
// (x) Sort an array in ascending order
// (xi)Sort an array in descending order

package main

import (
	"fmt"
)

func PrintArray(array []int) {
	for i, v := range array {
		fmt.Printf("Value at array[%d] is %d \n", i, v)
	}
}

func main() {

	var size int
	fmt.Println("Enter the size of your array:")
	fmt.Scanln(&size)

	array := make([]int, size) // array initiliazation

	// insert all the elements into array by taking input from user
	for i := 0; i < size; i++ {
		fmt.Printf("Enter value of array[%i] = ", i)
		fmt.Scanln(&array[i])
	}

	fmt.Println("----------MENU-----------")
	fmt.Println("1. Prints array element")
	fmt.Println("2. Length of an array")
	fmt.Println("3. Display array in reverse order")
	fmt.Println("4. Maximum element from an array ({1,2,3,6,5} ==> 6)")
	fmt.Println("5. Minimum element from an array")
	fmt.Println("6. Sum of all elements of an array")
	fmt.Println("7. Copy array elements to other array")
	fmt.Println("8. Print duplicate elements from an array")
	fmt.Println("9. Unique element from an array")
	fmt.Println("10. Sort an array in ascending order")
	fmt.Println("11. Sort an array in descending order")

	var choice int
	fmt.Println("Enter your choice:")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		PrintArray(array)
		return

	case 2:
	case 3:
	case 4:
	case 5:
	case 6:
	case 7:
	case 8:
	case 9:
	case 10:
	case 11:
	}

}
