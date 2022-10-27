package main
import (
	"fmt"
)

func main(){
	var n int
	fmt.Println("enter your no.")
	fmt.Scanln(&n)

		counter:=0

	for i:=1;i<=n/2;i++{
		if n%i==0{
		counter++
	}
	}
	if counter>1{
		fmt.Println("it is not a prime number")
	}else{
		fmt.Println("it is a prime number")
	}
}