// 1. Program (*USE: Functions, Switch cases and it should be menu based)
// (i)  Prints array element
// (ii) Length of an array
// (iii) Display array in reverse order
// (iv) Maximum element from an array ({1,2,3,6,5} ==> 6)
// (v) Minimum element from an array
// (vi) Sum of all elements of an array
// (vii) Copy array elements to other array
// (viii) Print duplicate elements from an array
// (ix) Unique element from an array
// (x) Sort an array in ascending order
// (xi) Sort an array in descending order

package main

import (
	"fmt"
)

func PrintArray(array []int) {
	for i, v := range array {
		fmt.Printf("Value at array[%d] is %d\n", i, v)
	}
}

func Length(array []int) int {
	return len(array)
}

func ReverseArray(array []int) {
	for i := len(array) - 1; i >= 0; i-- {
		fmt.Printf("value of array[%d] is %d\n", i, array[i])
	}
}

func MaxElement(array []int) int {
	if len(array) == 0 {
		// array is empty.. lets return 0 for now
		return 0
	}

	max := array[0]
	for i := 1; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
	}
	return max
}

func MinElement(array []int) int {
	if len(array) == 0 {
		// array is empty.. lets return 0 for now
		return 0
	}

	min := array[0]
	for i := 1; i < len(array); i++ {
		if array[i] < min {
			min = array[i]
		}
	}
	return min
}

func SumArray(array []int) int {
	sum := 0
	for _, v := range array {
		sum = sum + v
	}
	return sum
}

func AscendingArray(array []int) {

	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j +1] = temp
			}
		}
	}

	fmt.Println("Sorted array:", array)
}

func main() {

	var size int
	fmt.Println("Enter the size of your array:")
	fmt.Scanln(&size)

	if size == 0 {
		fmt.Println("No action can be performed!!")
		return
	}

	array := make([]int, size) // array initiliazation

	// insert all the elements into array by taking input from user
	for i := 0; i < size; i++ {
		fmt.Printf("Enter value of array[%d] = ", i)
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

	case 2:
		fmt.Println("length of array is :", Length(array))

	case 3:
		ReverseArray(array)

	case 4:
		fmt.Println("max element is:", MaxElement(array))

	case 5:
		fmt.Println("min element is:", MinElement(array))

	case 6:
		fmt.Println("sum of elements of array is:", SumArray(array))

	case 7:
		array2 := make([]int, size)
		copy(array2, array)
		fmt.Println("array2:", array2)

	case 8:
	case 9:

	case 10:
		AscendingArray(array)

	case 11:
	}

}