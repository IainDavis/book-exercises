package main

import "fmt"

func ex03() {
	var (
		b      byte   = 255
		smallI int32  = 2147483647
		bigI   uint64 = 18_446_744_073_709_551_615
	)

	fmt.Printf("\tbyte (pre-increment): %d\n", b)
	fmt.Printf("\tint32 (pre-inrement): %d\n", smallI)
	fmt.Printf("\tuint64 (pre-increment): %d\n", bigI)

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println("\t-----------")
	fmt.Printf("\tbyte (post-increment): %d\n", b)
	fmt.Printf("\tint32 (post-inrement): %d\n", smallI)
	fmt.Printf("\tuint64 (post-increment): %d\n", bigI)
	fmt.Println("")
	fmt.Println("Conclusion: bit overflow wraps around to the smallest allowable number (0 for unsigned)")
}
