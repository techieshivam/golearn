package main
import (
	"fmt"
)
func main(){
	
	for i:=1;i<=4;i++{
		for j:=3;j>=0;j--{
			if j>=i{
			fmt.Print(" ")
		}else{
			fmt.Print(i," ")
		
		}
		}
		fmt.Println()
	}
}