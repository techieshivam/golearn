package main
import (
	"fmt"
)

func max3(a,b,c int) int {
	if a>b && a>c{
		return a 
	}else if a<b && b>c{
		return b 
	}
	return c 
}
func main(){
    var a,b,c int

    fmt.Println("enter a no.")
    fmt.Scanln(&a)
	fmt.Println("enter a 2nd no.")
	fmt.Scanln(&b)
    fmt.Println("enter a 3rd no.")
    fmt.Scanln(&c)	


maximum:=max3(a,b,c)
fmt.Println("greatest bw 3 is",maximum)
}