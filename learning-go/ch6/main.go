package main

import (
	"fmt"

	"iaindavis.dev/learning/learning-go/ch6/exercises"
)

const showExamples = false

func main() {
	fmt.Println("Learning Go: Chapter 6")

	if showExamples {
		fmt.Println("Examples from the chapter text...")
		fmt.Println("\t", "-------------------")
	}

	fmt.Println("\t", "=== Exercises ===")

	fmt.Println("\t", "--- Exercise 01 ---")
	exercises.Exercise_6_1()
	fmt.Println("\t", "-------------------")

	fmt.Println("\t", "--- Exercise 02 ---")
	exercises.Exercise_6_2()
	fmt.Println("\t", "-------------------")
}
