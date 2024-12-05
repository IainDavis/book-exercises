package main

import "fmt"

func ex_2_3() {
	var x int = 10
	var b byte = 100
	var sum3 int = x + int(b)
	var sum4 byte = byte(x) + b

	fmt.Printf("\tx (int): %d;\n\t x (byte): %d\n", x, b)
	fmt.Printf("\tsum with integer conversion: %d\n", sum3)
	fmt.Printf("\tsum with byte conversion: %d\n", sum4)
}

func ex_2_3a() {
	var x int = 1000 // 1111101000 -> 1110 1000 (with truncation) = 232
	var b byte = 100 // 0001100100

	//   11
	//   1110 1000
	//   0110 0100
	// -------------------
	// 1 0100 1100 = 332
	// ^ bit overflow

	// truncate to 8 bits: 0100 1100 = 76

	var sum3 int = x + int(b)
	var sum4 byte = byte(x) + b

	fmt.Printf("\tx (int): %d;\n\t x (byte): %d\n", x, b)
	fmt.Printf("\tsum with integer conversion: %d\n", sum3) // = 1100
	fmt.Printf("\tsum with byte conversion: %d\n", sum4)    // = 76

	fmt.Println("\tint conversion to byte truncates most significant binary digits, until it fits in 8-bit width")
}
