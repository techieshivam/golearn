package main 

import (
	"fmt"
)

func main () {

	array:=[]int{4,2,5,8,9,7}
	var evenArray,oddArray []int 

	for _,v:=range array {
		if v%2==0{
			evenArray=append(evenArray,v)
		}else{
			oddArray=append(oddArray,v)
		}
	}
		fmt.Println(evenArray)
		fmt.Println(oddArray)
	}
