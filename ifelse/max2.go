package main
import (
	"fmt"
)


func maxi(a,b int) int {
	if a>b {
		return a
} else {
	return b
}

}


func main (){
	var a,b int 
	fmt.Println("enter no.")
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	maximum:=maxi(a,b)
	fmt.Println("bigger no. is",maximum)

}