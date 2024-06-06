package statistics

import "testing"

func TestStdDev(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5}
	var expected float64 = 1.4142135623730951
	output := StdDev(input)
	if output != expected {
		t.Errorf(`got %v, want %v`, output, expected)
	}
}
