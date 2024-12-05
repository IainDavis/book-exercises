package examples

import "fmt"

func Example01() {
	fmt.Println("\tExample 01 | Understanding Capacity")
	var x []int
	fmt.Println("\t\t\tUninitialized slice:")
	fmt.Println("\t\t\tx, len(x), cap(x):", x, len(x), cap(x))

	x = append(x, 10)
	fmt.Println("\t\tappend element: 10")
	fmt.Println("\t\t\tx, len(x), cap(x):", x, len(x), cap(x))

	fmt.Println("\t\tappend element: 20")
	x = append(x, 20)
	fmt.Println("\t\t\tx, len(x), cap(x):", x, len(x), cap(x))

	fmt.Println("\t\tappend element: 30")
	x = append(x, 30)
	fmt.Println("\t\t\tx, len(x), cap(x):", x, len(x), cap(x))

	fmt.Println("\t\tappend element: 40")
	x = append(x, 40)
	fmt.Println("\t\t\tx, len(x), cap(x):", x, len(x), cap(x))

	fmt.Println("\t\tappend element: 50")
	x = append(x, 50)
	fmt.Println("\t\t\tx, len(x), cap(x):", x, len(x), cap(x))

	fmt.Println("\n\t\t notice doubling of capacity each time the previous capacity is exhausted")
}
