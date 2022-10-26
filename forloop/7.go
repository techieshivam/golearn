package main

import (
	"fmt"
) 
	
func main(){
	var n,i,square int 
	fmt.Println("square of first N ")
	fmt.Scanln(&n)

	for i=0;i<=n;i++ {
	 square =i*i
	fmt.Println("square of",i,"is",square) 
	}

}
	
