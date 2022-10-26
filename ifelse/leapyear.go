package main
import (
	"fmt"
)

func leapyr(a int)string{
	if a%4==0 && a%100 !=0 {
		return "leap year"
	}
	return "not a leap year"
}


func main(){
 var a int
 fmt.Println("enter a year")
 fmt.Scanln(&a)

 isit:=leapyr(a)
 fmt.Println(isit)
}