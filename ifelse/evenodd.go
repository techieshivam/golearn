package main 
import (
	"fmt"
)

func evenodd(a int)string{
	if a%2==0{
		return "even"
	}
	return "odd"
}




func main(){
var a int
fmt.Println("enter a no.")
fmt.Scanln(&a)

isit:=evenodd(a)
fmt.Println("no. is",isit)
}