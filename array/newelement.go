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

fmt.Println("At which position you want to insert element:")
var position int 
fmt.Scanln(&position)

array = append(array,0) 
fmt.Println(array)
copy(array[position+1:], array[position:]) 
fmt.Println(array)
array[position] = 99

fmt.Println(array)

}





