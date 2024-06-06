package statistics

import (
	"reflect"
	"testing"
)

func TestAverage(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5}
	output := Average(input)
	expected := 3
	if output != float64(expected) {
		t.Errorf(`Got %v and want %v`, output, input)
	}
}

func TestStringConverter(t *testing.T) {
	strSlice := []string{"1", "2", "3", "4", "5"}
	fltSlice := []float64{1, 2, 3, 4, 5}
	output := StringConverter(strSlice)
	if !reflect.DeepEqual(output, fltSlice) {
		t.Errorf(`Got %v and want %v`, output, fltSlice)
	}
}
