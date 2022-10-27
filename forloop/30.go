package main  

import(
	"fmt"
)

func main(){
	var n,r int 
	fmt.Println("enter your no.")
	fmt.Scanln(&n)
	sum:=0
	for n!=0 {
		r=n%10
		sum=sum*10+r
		n=n/10
	}
	fmt.Println(sum)
}