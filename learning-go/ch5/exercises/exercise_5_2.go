package exercises

import (
	"fmt"
	"os"
)

func fileLen(s string) (int, error) {
	buffer := make([]byte, 2024)

	f, err := os.Open(s)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	count, err := f.Read(buffer)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func Exercise_5_2() {
	msg, err := fileLen("/Users/iain/Study/book-exercises/learning-go/ch5/exercises/exercise_5_2.go")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("file length:", msg)

}
