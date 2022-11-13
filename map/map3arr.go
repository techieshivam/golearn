package main

import (
	"fmt"
)

type student struct {
	roll   int
	name   string
	class  int
	gender string
}

func linearSearch(roll int, students []student) student {
	for _, v := range students {
		if v.roll == roll {
			return v
		}
	}
	return student{}
}

func main() {
	students := make([]student, 3)
	students[0] = student{
		roll:   1,
		name:   "shivam",
		class:  10,
		gender: "M",
	}
	students[1] = student{
		roll:   2,
		name:   "satyam",
		class:  10,
		gender: "M",
	}
	students[2] = student{
		roll:   3,
		name:   "naman",
		class:  10,
		gender: "M",
	}
	fmt.Println("student list is", students)
	x := linearSearch(1, students)
	fmt.Println(x)
}
