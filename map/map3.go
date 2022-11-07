package main

import "fmt"

type Student struct {
	roll   int
	name   string
	class  int
	gender string
	age    int
}

func main() {

	stud1 := Student{
		roll:   1,
		name:   "shivam",
		class:  10,
		gender: "M",
		age:    18,
	}
	stud2 := Student{
		roll:   2,
		name:   "satyam",
		class:  10,
		gender: "M",
		age:    18,
	}
	stud3 := Student{
		roll:   3,
		name:   "parikshit",
		class:  10,
		gender: "M",
		age:    18,
	}
	stud4 := Student{
		roll:   4,
		name:   "rahul",
		class:  10,
		gender: "M",
		age:    18,
	}
	stud5 := Student{
		roll:   5,
		name:   "deepanshu",
		class:  10,
		gender: "M",
		age:    18,
	}
	stud6 := Student{
		roll:   6,
		name:   "sahil",
		class:  10,
		gender: "M",
		age:    18,
	}
	stud7 := Student{
		roll:   7,
		name:   "mahesh",
		class:  10,
		gender: "M",
		age:    18,
	}
	stud8 := Student{
		roll:   8,
		name:   "naman",
		class:  10,
		gender: "M",
		age:    18,
	}

	var roll_number int
	fmt.Println("These are the roll_number in class")
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
	fmt.Println("5")
	fmt.Println("6")
	fmt.Println("7")
	fmt.Println("8")

	fmt.Println("record of which student you want")
	fmt.Scanln(&roll_number)

	students := map[int]Student{
		1: stud1,
		2: stud2,
		3: stud3,
		4: stud4,
		5: stud5,
		6: stud6,
		7: stud7,
		8: stud8,
	}

	student_info := students[roll_number]

	fmt.Println(student_info)

}
