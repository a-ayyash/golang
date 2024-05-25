package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected %d but got %d", sum, expected)
	}
}

// the comment below is important to run the example
func ExampleAdd() {
	sum := Add(5, 1)
	fmt.Println(sum)
	// Output: 6
}

func Add(x, y int) int {
	return x + y
}
