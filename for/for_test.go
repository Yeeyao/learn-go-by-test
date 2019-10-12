package main

import "testing"

func TestFor(t *testing.T) {

	t.Run("for test", func(t *testing.T) {
		repeated := For(10)
		expected := "aaaaaaaaaa"
		if repeated != expected {
			t.Errorf("expected '%s' but got '%s'", expected, repeated)
		}
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		For(10)
	}
}
