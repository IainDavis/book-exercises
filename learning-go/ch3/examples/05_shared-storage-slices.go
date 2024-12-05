package examples

import "fmt"

func Example05() {

	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:]
	fmt.Println(`
		x := []string{"a", "b", "c", "d"}
		y := x[:2]
		z := x[1:]
	`)

	fmt.Println("\t", "Before transformations:")
	fmt.Println("\t\t", "x:", x)
	fmt.Println("\t\t", "y:", y)
	fmt.Println("\t\t", "z:", z)

	x[1] = "y"
	y[0] = "x"
	z[1] = "z"

	fmt.Println("\tTransformations:")
	fmt.Println(`
		x[1] = "y"
		y[0] = "x"
		z[1] = "z"`)
	fmt.Println()

	fmt.Println("\t", "After transformations:")
	fmt.Println("\t\t", "x:", x)
	fmt.Println("\t\t", "y:", y)
	fmt.Println("\t\t", "z:", z)
}
