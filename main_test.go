// main_test.go contains unit tests and benchmarks for the Coeff function
package main

import (
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

func round(x float64) float64 {
	return math.Round(x*10000) / 10000
}

// TestCoefficients tests the Coeff function against known values
func TestCoefficients(t *testing.T) {
	tests := []struct {
		name        string
		points      []stats.Coordinate
		actualint   float64
		actualslope float64
	}{
		{"Set 1", c1, 3.0001, 0.5001},
		{"Set 2", c2, 3.0009, 0.5000},
		{"Set 3", c3, 3.0025, 0.4997},
		{"Set 4", c4, 3.0017, 0.4999},
	}
	// Loop through each test case and compare the results
	for _, test := range tests {
		resIntercept, resSlope := Coeff(test.points)

		if round(resIntercept) != round(test.actualint) {
			t.Errorf("%s intercept = %.4f, want %.4f",
				test.name, resIntercept, test.actualint)
		}

		if round(resSlope) != round(test.actualslope) {
			t.Errorf("%s slope = %.4f, want %.4f",
				test.name, resSlope, test.actualslope)
		}
	}
}

// BenchmarkCoefficients benchmarks the Coeff function to measure its performance
func BenchmarkCoefficients(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Coeff(c1)
		Coeff(c2)
		Coeff(c3)
		Coeff(c4)
	}
}
