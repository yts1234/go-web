package iteration

import "testing"

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("expected %q but got %q", got, want)
		}
	}
	t.Run("Repeat word printing by repeated n count supplied, default 5", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := "aaaaa"

		assertCorrectMessage(t, repeated, expected)
	})
	t.Run("Repeat word printing by repeated count supplied, repeat count = 10", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"

		assertCorrectMessage(t, repeated, expected)
	})
}

// func BenchmarkRepeat(b *testing.B) {
// 	Repeat("a", 0)
// }
