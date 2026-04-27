// This program validates the montanaflynn/stats Go package by performing
// linear regression and comparing results to Python and R
package main

import (
	"fmt"
	"math"
	"runtime"
	"time"

	"github.com/montanaflynn/stats"
	"gonum.org/v1/gonum/stat/distuv"
)

// Anscombe's quartet data points
var c1 = []stats.Coordinate{
	{X: 10, Y: 8.04},
	{X: 8, Y: 6.95},
	{X: 13, Y: 7.58},
	{X: 9, Y: 8.81},
	{X: 11, Y: 8.33},
	{X: 14, Y: 9.96},
	{X: 6, Y: 7.24},
	{X: 4, Y: 4.26},
	{X: 12, Y: 10.84},
	{X: 7, Y: 4.82},
	{X: 5, Y: 5.68},
}

var c2 = []stats.Coordinate{
	{X: 10, Y: 9.14},
	{X: 8, Y: 8.14},
	{X: 13, Y: 8.74},
	{X: 9, Y: 8.77},
	{X: 11, Y: 9.26},
	{X: 14, Y: 8.10},
	{X: 6, Y: 6.13},
	{X: 4, Y: 3.10},
	{X: 12, Y: 9.13},
	{X: 7, Y: 7.26},
	{X: 5, Y: 4.74},
}

var c3 = []stats.Coordinate{
	{X: 10, Y: 7.46},
	{X: 8, Y: 6.77},
	{X: 13, Y: 12.74},
	{X: 9, Y: 7.11},
	{X: 11, Y: 7.81},
	{X: 14, Y: 8.84},
	{X: 6, Y: 6.08},
	{X: 4, Y: 5.39},
	{X: 12, Y: 8.15},
	{X: 7, Y: 6.42},
	{X: 5, Y: 5.73},
}

var c4 = []stats.Coordinate{
	{X: 8, Y: 6.58},
	{X: 8, Y: 5.76},
	{X: 8, Y: 7.71},
	{X: 8, Y: 8.84},
	{X: 8, Y: 8.47},
	{X: 8, Y: 7.04},
	{X: 8, Y: 5.25},
	{X: 19, Y: 12.50},
	{X: 8, Y: 5.56},
	{X: 8, Y: 7.91},
	{X: 8, Y: 6.89},
}

// Coeff calculates the slope and intercept
func Coeff(points []stats.Coordinate) (float64, float64) {
	r, _ := stats.LinearRegression(points)

	p1 := r[0]
	p2 := r[len(r)-1]

	if p1.X == p2.X {
		for a := 0; a < len(r); a++ {
			if r[a].X != p1.X {
				p2 = r[a]
				break
			}
		}
	}

	slope := (p2.Y - p1.Y) / (p2.X - p1.X)
	intercept := p1.Y - slope*p1.X

	return intercept, slope
}

// regression performs linear regression and calculates statistics
func regression(name string, points []stats.Coordinate) {
	intercept, slope := Coeff(points)
	// calculate R-squared
	var sumY float64
	for _, p := range points {
		sumY += p.Y
	}
	meanY := sumY / float64(len(points))
	// calculate RSS and TSS
	var rss, sst float64
	for _, p := range points {
		predicted := intercept + slope*p.X
		rss += (p.Y - predicted) * (p.Y - predicted)
		sst += (p.Y - meanY) * (p.Y - meanY)
	}
	r2 := 1.0 - (rss / sst)

	var sumX float64
	for _, p := range points {
		sumX += p.X
	}
	meanX := sumX / float64(len(points))
	var sumSqDiffX float64
	for _, p := range points {
		sumSqDiffX += (p.X - meanX) * (p.X - meanX)
	}
	mse := rss / float64(len(points)-2)
	// standard error
	se := math.Sqrt(mse / sumSqDiffX)
	// t-statistic
	tStat := slope / se
	// p-value and F-statistic
	df := float64(len(points) - 2)
	tDist := distuv.StudentsT{Mu: 0, Sigma: 1, Nu: df}
	pValue := 2 * (1 - tDist.CDF(math.Abs(tStat)))
	fStat := ((sst - rss) / 1) / (rss / float64(len(points)-2))
	// adjusted R-squared
	n := float64(len(points))
	ar := 1 - ((1 - r2) * (n - 1) / (n - 2))
	// print results
	fmt.Printf("%s: intercept = %.6f slope = %.6f R-squared = %.4f SE = %.4f t = %.4f p = %.4f F = %.4f Adjusted R-squared = %.4f\n", name, intercept, slope, r2, se, tStat, pValue, fStat, ar)
}

// execute the regression analysis
// measure execution time and memory usage
func main() {

	var ms, me runtime.MemStats
	runtime.GC() // force garbage collection
	// start memory and time measurement
	runtime.ReadMemStats(&ms)
	start := time.Now()

	regression("Set 1", c1)
	regression("Set 2", c2)
	regression("Set 3", c3)
	regression("Set 4", c4)

	// end time and memory measurement
	elapsed := time.Since(start)
	runtime.ReadMemStats(&me)
	// print results
	fmt.Printf("\nExecution Time: %v seconds\n", elapsed.Seconds())
	fmt.Printf("Memory: %v KB\n", (me.Alloc-ms.Alloc)/1024)

}
