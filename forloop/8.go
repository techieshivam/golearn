package main

import (
	"fmt"
) 
	
func main() {
	var n,i,cube int 
	fmt.Println("cube of first N ")
	fmt.Scanln(&n)

	for i=0;i<=n;i++ {
	 cube =i*i*i
	fmt.Println("cube of",i,"is",cube) 
	}

}