package main
import (
	"fmt"
)
func main(){
	counter:=1
	for i:=4;i>0;i--{
		for j:=0;j<=4;j++{
			if j<i{
			fmt.Print(" ")
		}else{
			fmt.Print(counter," ")
			counter++
		}
		}
		fmt.Println()
	}
}