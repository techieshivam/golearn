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

	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j] < array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j +1] = temp
			}
		}
	}	

	temp1:=array[0]
	array[0]=array[1]
	array[1]=temp1


	fmt.Println("second largest on first place",array)





}