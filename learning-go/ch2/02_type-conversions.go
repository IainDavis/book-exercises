package main

import "fmt"

func ex_2_2() {
	var x int = 10
	var y float64 = 30.2
	var sum1 float64 = float64(x) + y
	var sum2 int = x + int(y)

	fmt.Printf("\tx (int) = %d; y (float64) = %f\n", x, y)
	fmt.Printf("\tsum with float cast: %f\n", sum1)
	fmt.Printf("\tsum with integer cast: %d\n", sum2)
}
