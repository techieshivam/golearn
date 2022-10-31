package main

import (
	"fmt"
)

func PrintElement(array []int) {
	for i, v := range array {
		fmt.Printf("Value at array[%d] is %d\n", i, v)
	}
}

func LengthArray(array []int) int {
	return len(array)
}

func ReverseArray(array []int) {
	for i := len(array) - 1; i >= 0; i-- {
		fmt.Printf("value of array[%d] is %d\n", i, array[i])
	}
}

func Maxi(array []int) int {
	if len(array) == 0 {
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

func Mini(array []int) int {
	if len(array) == 0 {
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

func CopyElements() {

}

func Duplicates() {

}

func Unique() {

}

func AscendingArray(array []int) {

	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}

	fmt.Println("Sorted array:", array)
}

func decendingArray(array []int) {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j] < array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
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
		fmt.Println("No action can be performed!")
		return
	}

	array := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Printf("Enter value of array[%d] = ", i)
		fmt.Scanln(&array[i])
	}

	var option int

	fmt.Println("what you want to do:")
	fmt.Println("1: Prints array element")
	fmt.Println("2: Length of an array")
	fmt.Println("3: Display array in reverse order")
	fmt.Println("4: Maximum element from an array")
	fmt.Println("5: Minimum element from an array")
	fmt.Println("6: Sum of all elements of an array")
	fmt.Println("7: Copy array elements to other array")
	fmt.Println("8: Print duplicate elements from an array")
	fmt.Println("9: Unique element from an array")
	fmt.Println("10: Sort an array in ascending order")
	fmt.Println("11: Sort an array in descending order")
	fmt.Println("Enter option: ")
	fmt.Scanln(&option)
	fmt.Println("your choice is", option)

	switch option {

	case 1:
		PrintElement(array)

	case 2:
		fmt.Println("length of array is :", LengthArray(array))

	case 3:
		ReverseArray(array)

	case 4:
		fmt.Println("max element is:", Maxi(array))

	case 5:
		fmt.Println("max element is:", Mini(array))

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
		decendingArray(array)

	}

}
