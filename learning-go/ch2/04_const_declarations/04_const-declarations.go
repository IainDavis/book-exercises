package main

import "fmt"

const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10

func ex_2_4() {
	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)

	x = x + 1 // will not compile
	y = "bye" // will not compile
	fmt.Println(x)
	fmt.Println(y)
}
