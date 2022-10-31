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

	array:=make([]int,size)
	
	for i:=0;i<size;i++{
		fmt.Printf("Enter value of array[%d] = ", i)
		fmt.Scanln(&array[i])
		}

	fmt.Println("Array is :",array)

	fmt.Println("which element you want to delete")
	var position int
	fmt.Scanln(&position)

	array=append(array[: position],array[position+1:]...)

	fmt.Println("New Array is :",array)




}