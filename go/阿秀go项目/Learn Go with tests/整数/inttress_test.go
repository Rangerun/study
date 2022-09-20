package intress

import "testing"


func add(a, b int) int {
	return a + b
}

func TestInt(t *testing.T) {
	sum := add(1, 2)
	expected := 4
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
