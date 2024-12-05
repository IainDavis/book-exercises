package main

import "fmt"

func ex02() {
	const value = 10
	var i int = value
	var f float64 = value

	fmt.Printf("\tconst value: %d\n", value)
	fmt.Printf("\tvalue as integer: %d\n", i)
	fmt.Printf("\tvalue as floating-point number: %f\n", f)
}
