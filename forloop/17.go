package main
import (
	"fmt"
)
func main(){
	for i:=5;i>0;i--{
		for j:=0;j<=5;j++{
			if j<i{
				fmt.Print(" ")
			}else{
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
}