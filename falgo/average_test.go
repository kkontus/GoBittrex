package falgo_test

import (
	. "GoBittrex/falgo"
	"testing"
	"reflect"
)

func TestAverage(t *testing.T) {
	// test data
	test := []float64{82.5, 148, 143, 127, 132, 133}

	// test data, we will take this from getTicks
	l := []float64{50, 98, 93, 77, 82, 83}
	h := []float64{115, 198, 193, 177, 182, 183}

	avg := Average(l, h)
	if reflect.DeepEqual(test, avg) {
		t.Log("test == avg")
	}

	if !reflect.DeepEqual(test, avg) {
		t.Error("test != avg")
	}
}
