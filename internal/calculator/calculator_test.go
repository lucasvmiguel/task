package calculator

import "testing"

func TestSum(t *testing.T) {
	expected := 5
	got := Sum(2, 3)

	if got != expected {
		t.Errorf("sum should be %d but got %d", expected, got)
	}
}
