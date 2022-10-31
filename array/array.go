package main

import (
	"fmt"
)

func main () {
	var a[5] int 
	a[0]=7
	a[1]=2
	a[2]=3
    a[3]=6
    a[4]=4
    fmt.Println(a)

    var veglist=[...] string{"tomatato","potato","rddish"}
    fmt.Println(veglist)
 fmt.Println(len(veglist))

 a3:=a
 fmt.Println(a3)



 fmt.Println(a[2:4])

}