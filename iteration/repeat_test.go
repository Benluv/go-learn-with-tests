package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Expected '%q' to be '%q'", repeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func TestCountRepeat(t *testing.T) {
	t.Run("repeat a character five times and return the total amount of appearances for said character", func(t *testing.T) {
		counted := CountRepeat("a", 5)
		expected := 5
		assertCorrectMessage(t, counted, expected)
	})

	t.Run("return the number 8 when count 0 value is provided", func(t *testing.T) {
		counted := CountRepeat("b", 0)
		expected := 8
		assertCorrectMessage(t, counted, expected)
	})

}

func assertCorrectMessage(t testing.TB, counted, expected int) {
	t.Helper()
	if counted != expected {
		t.Errorf("Expected '%d' to be '%d'", counted, expected)
	}
}
