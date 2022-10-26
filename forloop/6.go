package main

import (
	"fmt"
) 
	
func main(){
	var n,i int 
	fmt.Println("avg of first N ")
	fmt.Scanln(&n)
	sum:=0
	avg:=0
	for i=1;i<=n;i++ {
	 sum=sum+i
	 avg=sum/i
	
	}
	fmt.Println("avg is",avg) 
}
	
