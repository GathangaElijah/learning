package statistics

import "testing"

func TestVariance(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5}
	expected := 2
	output := Variance(input)
	if output != float64(expected) {
		t.Errorf(`got %v want %v`, output, expected)
	}
}
