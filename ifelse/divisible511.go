package main 
import (
	"fmt"
)

func divi511( a int)string{
	if a%5==0 && a%11==0 {
		return "divisible by 5 and 11"
	}
	return "not divisible by 5 and 11"
}


func main(){
	var a int
	fmt.Println("enter ano.")
fmt.Scanln(&a)

weather:=divi511(a)
fmt.Println(weather)


}