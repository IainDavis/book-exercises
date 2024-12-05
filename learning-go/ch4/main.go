package main

import (
	"fmt"

	"iaindavis.dev/learning/learning-go/ch4/examples"
)

const showExamples = true

func main() {
	if showExamples {
		fmt.Println("Examples from Chapter 4 body:")
		fmt.Println("\t", "Example 4|12 -- FizzBuzz with `continue`")
		examples.Example4_12()
		fmt.Println("---------------")
	}
}
