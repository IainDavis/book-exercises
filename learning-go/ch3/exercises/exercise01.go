package exercises

import "fmt"

func Exercise01() {
	var greetings = []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	var x = greetings[:2]
	var y = greetings[1:4]
	var z = greetings[4:]

	fmt.Println("\t\t", "greetings:", greetings)
	fmt.Println("\t\t", "x:", x)
	fmt.Println("\t\t", "y:", y)
	fmt.Println("\t\t", "z:", z)
}
