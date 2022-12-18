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

   for i, v := range array {
		fmt.Printf("Value at array[%d] is %d\n", i, v)
	}


        evenSize := 0;
        oddSize := 0;

        for i:=0;i<size;i++{
        	if array[i]%2==0{
        		evenSize++
        	}else {
        		oddSize++
        	}
        }

        fmt.Println(evenSize)
		fmt.Println(oddSize)


		
		evenArray:=make([]int,evenSize)
		oddArray:=make([]int,oddSize)
		j := 0 
		k := 0
        for i:= 0;i<size;i++ {
            if (array[i]%2 == 0){
                	evenArray[j] =array [i]
                	j++
            }else{
                oddArray[k] =array[i] 
                k++
            }
        }

        fmt.Println("Even Array contains: ",evenArray)
        fmt.Println("odd Array contains: ",oddArray)
}