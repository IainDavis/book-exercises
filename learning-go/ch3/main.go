package main

import (
	"fmt"

	"iaindavis.dev/learning/learning-go/ch3/examples"
	"iaindavis.dev/learning/learning-go/ch3/exercises"
)

const showExamples = false

func main() {
	if showExamples {
		fmt.Println("Learning Go: Chapter 3")
		fmt.Println("\tExamples from chapter body")
		examples.Example01()
		fmt.Println("-------------------")
		examples.Example04()
		fmt.Println("-------------------")
		examples.Example05()
		fmt.Println("-------------------")
	}

	fmt.Println("\tEnd-of-Chapter Exercises")
	exercises.Exercise01()
	fmt.Println("-------------------")
	exercises.Exercise02()
	fmt.Println("-------------------")
	exercises.Exercise03()
	fmt.Println("-------------------")
}
