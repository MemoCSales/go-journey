package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

func Repeat(c string) (repeat_output string) {
	for i := 0; i < 5; i++ {
		repeat_output += c
	}
	return
}
