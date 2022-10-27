package main

import (
	"fmt"
) 
	
func main(){
	var n,i  int 
	fmt.Println("square of first N ")
	fmt.Scanln(&n)
	sum:=0
	square:=0

	for i=1;i<=n;i++ {
	 square =i*i
	
	 sum=sum+square

fmt.Println("sum of square first",i,"is ",sum) 
	
	}
	fmt.Println("sum of square is",n,sum) 

}