package main 
import (
	"fmt"
)

func posinegi(a int) string {
	if a>0{
		return "no is positive"
	}else if a<0{
		return "no. is negative"
	}
	return "zero"
}


func main(){
var a int 
fmt.Println("eneter a no.")
fmt.Scanln(&a)

isit:=posinegi(a)
fmt.Println(isit)
}