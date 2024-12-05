package main

import "fmt"

func Example4_12() {
	for i, hit := 1, false; i <= 100; i, hit = i+1, false {
		fmt.Print("\t\t")
		if i%3 == 0 {
			hit = true
			fmt.Print("Fizz")
		}
		if i%5 == 0 {
			hit = true
			fmt.Print("Buzz")
		}
		if hit {
			fmt.Println()
			continue
		}
		fmt.Println(i)
	}
}
