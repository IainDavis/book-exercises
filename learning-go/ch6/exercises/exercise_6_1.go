package exercises

import "fmt"

/*
Create a struct named Person with three fields: FirstName and
LastName of type string, and Age of type int. Write a function
called MakePerson that takes in firstName, lastName, and age and
returns a Person. Write a second function MakePersonPointer that
takes in firstName, lastName, and age and returns a *Person. Call
both from `main`. Compile your program with go build
-gcflags="-m". This both compiles your code and prints out which
values escape to the heap. Are you surprised about what escapes?
*/
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName string, lastName string, age int) Person {
	person := Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
	return person
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {

	person := Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}

	return &person
}

func Exercise_6_1() {
	personValue := MakePerson("Maynard", "Krebbs", 32)

	fmt.Println("person value:", personValue)

	personPointer := MakePersonPointer("Maynard", "Krebbs", 32)
	fmt.Println("person pointer:", personPointer)
}

/*
	It looks like all strings and both Person structs are moved to the
	heap. Not really sure what to make of that. I'd expect the structs
	to be on the stack, since they are of known size.
*/
