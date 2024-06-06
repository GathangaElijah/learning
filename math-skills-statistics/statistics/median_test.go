package statistics

import "testing"

func TestMedian(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5}
	var expected float64 = 3
	output := Median(input)
	if output != expected {
		t.Errorf(`got %v, want %v`, output, expected)
	}
}
