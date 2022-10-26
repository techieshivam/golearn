package main
import (
	"fmt"
)
func main(){
var n int 
fmt.Println("enter the no.")
fmt.Scanln(&n)
factorial:=1
for i:=n;i>=1;i--{
	factorial=factorial*i
	}
	fmt.Print("factorial of no. is:",factorial )
}