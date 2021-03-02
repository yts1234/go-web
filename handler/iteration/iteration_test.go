package iteration

import "testing"

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t )
	t.Run("Repeat word printing by repeated n count supplied, default 5", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
	t.Run("Repeat word printing by repeated count supplied, repeat count = 10", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

// func BenchmarkRepeat(b *testing.B) {
// 	Repeat("a", 0)
// }
