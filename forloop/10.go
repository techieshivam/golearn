package main
import (
	"fmt"
)
func main(){
	var n int 
	fmt.Println("enter the no.")
	fmt.Scanln(&n)
	for i:=1;i<=10;i++{
		table:=i*n
		fmt.Println(n,"*",i,"=",table)
	}
}