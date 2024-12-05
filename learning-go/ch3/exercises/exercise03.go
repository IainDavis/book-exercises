package exercises

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func Exercise03() {
	first := Employee{
		"Joe-Bob",
		"Briggs",
		1,
	}

	second := Employee{
		id:        2,
		firstName: "Rhonda",
		lastName:  "Shearer",
	}

	var third Employee
	third.firstName = "Gilbert"
	third.lastName = "Gottfried"
	third.id = 3

	fmt.Println("first:", first)
	fmt.Println("second:", second)
	fmt.Println("third:", third)
}
