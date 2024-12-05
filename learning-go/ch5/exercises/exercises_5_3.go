package exercises

import "fmt"

func prefixer(p string) func(string) string {
	return func(s string) string { return p + " " + s }
}

func Exercise_5_3() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Maria"))
}
