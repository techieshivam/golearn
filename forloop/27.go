package main
import (
	"fmt"
)

func main(){
	var n1,n2,i,j int

	fmt.Println("enter your smaller no.")
	fmt.Scanln(&n1)
	fmt.Println("enter greater no. than first no")
	fmt.Scanln(&n2)

		counter:=0

	for i=1;i<=n2;i++{
		for j=n1+1;j<n2;j++{
			if j%i==0 {
				counter++ 
					}
		if j==1{
			fmt.Println(j)
		}
}


	}
	
}