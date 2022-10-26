package main
import (
	"fmt"
)
func main (){
	var n int 
	fmt.Println("enter a no")
	fmt.Scanln(&n)
	for i:=1;i<=n;i++{
		fmt.Println(i)
	}
}