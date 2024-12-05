package exercises

import (
	"fmt"

	common "iaindavis.dev/learning/learning-go/ch8/learning-go/ch8/common"
)

/*
Write a generic function that doubles the value of any integer or
float that's passed in to it. Define any needed generic
interfaces.
*/

func Double[N common.Number](n N) N { return n * 2 }

func Exercise_8_1() {
	fmt.Println("doubled int:", Double(10))
	fmt.Println("doubled float:", Double(30.8))
}
