package main

import (
	"fmt"
) 
	
func main(){
	var n,i int 
	fmt.Println("sum of first N ")
	fmt.Scanln(&n)
	sum:=0
	for i=0;i<=n;i++ {
	 sum=sum+i
	
	}
	fmt.Println("sum is",sum) 
}
	

