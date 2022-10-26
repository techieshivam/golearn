package main
import(
	"fmt"
)

func profitloss( costprice,sellingprice int)(string,int){
	if costprice>sellingprice {
		return "you have a loss of",costprice-sellingprice 
	}else if costprice <sellingprice{
		return "you have a profit of",sellingprice-costprice
	}
	return "you dont have any profit or loss",0
}



func main(){
var a,b int 
fmt.Println("enter the cost price")
fmt.Scanln(&a)
fmt.Println("enter the selling price")
fmt.Scanln(&b)


isit,isit1:=profitloss(a,b)
fmt.Println(isit,isit1)
}