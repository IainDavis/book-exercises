package exercises

/*
	The simple calculator program doesn't handle one error case:
	division by zero. Change the function signature for the math
	operations to return both an `int` and an `error`. In the `div`
	function, if the divisor is 0, return `errors.New("division by
	zero")` for the error. In all other cases return `nil`. Adjust the
	`main` function to check for this error.
*/

import (
	"errors"
	"fmt"
	"strconv"
)

type opFuncType func(int, int) (int, error)

var opMap = map[string]opFuncType{
	"+": func(i int, j int) (int, error) { return i + j, nil },
	"-": func(i int, j int) (int, error) { return i - j, nil },
	"*": func(i, j int) (int, error) { return i * j, nil },
	"/": func(i, j int) (int, error) {
		if j == 0 {
			return 0, errors.New("division by zero")
		}
		return i / j, nil
	},
}

func Exercise_5_1() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
		{"10", "/", "0"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}

		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result, err := opFunc(p1, p2)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(result)
	}
}
