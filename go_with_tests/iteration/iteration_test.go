package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	t.Run("Ensures it repeats 5 times", func(t *testing.T) {
		got := Repeat("a")
		want := "aaaaa"

		if got != want {
			t.Errorf("got %s; want %s", got, want)
		}
	})
}

func Repeat(s string) string {
	str := ""
	for i := 0; i < 5; i++ {
		str += s
	}
	return str
}

// go test -bench=.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func ExampleRepeat() {
	str := Repeat("a")
	fmt.Println(str)

	// Output: aaaaa
}
