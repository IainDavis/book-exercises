package examples_test

import (
	"testing"

	"iaindavis.dev/learning/learning-go/ch8/learning-go/ch15/examples"
)

func TestAddNumbers(t *testing.T) {
	result := examples.AddNumbers(2, 3)
	if result != 5 {
		t.Error("incorrect result: expected 5, got", result)
	}
}
