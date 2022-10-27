package main
import (
	"fmt"
)
func main(){
	for i:=0;i<=5;i++{
		for j:=1;j<=5;j++{
			if j>i{
			fmt.Print("* ")
		}else{
			fmt.Print(" ")
		}
		}
		fmt.Println()
	}
}