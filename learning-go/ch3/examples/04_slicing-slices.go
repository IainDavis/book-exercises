package examples

import "fmt"

func announce() {
	fmt.Println("\tExample 04 | Understanding Capacity")
}

func Example04() {
	announce()

	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]

	fmt.Println("\tfull array:")
	fmt.Println("(\t\t", "x := []string{\"a\", \"b\", \"c\", \"d\"})")
	fmt.Println("\t\t", x)
	fmt.Println()

	fmt.Println("\tup to index 2 (excl):")
	fmt.Println("\t\t", "y := x[:2]")
	fmt.Println("\t\t", y)
	fmt.Println()

	fmt.Println("\tfrom index 1 (incl.) to the end:")
	fmt.Println("\t\t", "z := x[1:]")
	fmt.Println("\t\t", z)
	fmt.Println()

	fmt.Println("\tfrom index 1 (incl.) to three (excl.)")
	fmt.Println("\t\t", "d := x[1:3]")
	fmt.Println("\t\t", d)
	fmt.Println()

	fmt.Println("\tno start or end indices (new variable sharing the same memory, length, and capacity):")
	fmt.Println("\t\t", "e := x[:]")
	fmt.Println("\t\t", e)
}
