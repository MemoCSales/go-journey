package iteration

import (
	"fmt"
	"strings"
	"testing"
)

const repeatCount = 5

func TestRepeat(t *testing.T) {
	assertIteration := func(t testing.TB, got, want string) {
		t.Helper() 	// Marca esta función como una función auxiliar de prueba, no una prueba en sí misma
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("Repeating 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertIteration(t, repeated, expected)
	})

	t.Run("Repeating 0 times", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""
		assertIteration(t, repeated, expected)
	})
}

func Benchmark(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}

func Repeat(character string, count int) string {
	var repeated strings.Builder
	for i := 0; i < count; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}

func ExampleRepeat() {
	repeat := Repeat("a", 5)
	fmt.Println(repeat)
	// Output: aaaaa
}

func ExampleRepeat_zero() {
	repeat := Repeat("b", 0)
	fmt.Println(repeat)
	// Output: 
}

func ExampleRepeat_multipleCharacters() {
	repeat := Repeat("abc", 3)
	fmt.Println(repeat)
	// Output: abcabcabc
}
