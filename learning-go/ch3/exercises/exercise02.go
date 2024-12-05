package exercises

import "fmt"

func Exercise02() {
	greeting := "Hi 👩 and 👨"
	runes := []rune(greeting)
	fmt.Println(string(runes[3]))
}
