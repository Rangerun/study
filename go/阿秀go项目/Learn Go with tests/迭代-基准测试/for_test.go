package forTest

import "testing"

func repeat(a string) string {
	for i := 0; i < 5; i++ {
		a += a
	}
	return a
}

func TestforString(t *testing.T) {
	a := repeat("a")
	b := "aaaaa"
	if a != b {
		t.Errorf("a != b")
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeat("a")
	}
}