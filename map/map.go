package main

import (
	"fmt"
)

func main() {
	makeMap := make(map[int]int)
	makeMap[1] = 7
	makeMap[2] = 8
	makeMap[3] = 9
	makeMap[4] = 8
	makeMap[5] = 9
	makeMap[6] = 4

	var count int

	for _, makeMap := range makeMap {
		count += len(makeMap)
	}

	fmt.Println(makeMap)
	fmt.Println(count)
}
