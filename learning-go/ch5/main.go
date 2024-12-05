package main

import (
	"fmt"

	"iaindavis.dev/learning/learning-go/ch5/exercises"
)

const showExamples = false

func main() {
	fmt.Println("----- Chapter 5 Content -----")
	if showExamples {
		fmt.Println("\tChapter 5 executable examples")
		fmt.Println("\t-----------------")
	}

	fmt.Println("\t", "Chapter 5 Exercises")
	fmt.Println("\t", "*** Exercise 1 ***")
	exercises.Exercise_5_1()
	fmt.Println("\t", "==================")
	fmt.Println("\t", "*** Exercise 2 ***")
	exercises.Exercise_5_2()
	fmt.Println("\t", "==================")
	fmt.Println("\t", "*** Exercise 3 ***")
	exercises.Exercise_5_3()
	fmt.Println("\t", "==================")
}
