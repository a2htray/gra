package gra

import (
	"math"
)

type axis int

var Rho = 0.5

const (
	AxisRow axis = iota
	AxisColumn
)

// ComputeGRGs compute grey relational grades series by series
// By default, compute the grades along with row
func ComputeGRGs(values [][]float64, axis axis) [][]float64 {
	if axis == AxisColumn {
		return computeGRGs(t(values))
	} else {
		return computeGRGs(values)
	}
}

func computeGRGs(values [][]float64) [][]float64 {
	valuesInput := copyValuesUseCopy(values)
	r, _ := shape(valuesInput)

	res := make([][]float64, r)

	for i := 0; i < r; i++ {
		ref := valuesInput[i]
		valuesTemp := copyValuesUseCopy(valuesInput)
		valuesTemp = append(valuesTemp[:i], values[i+1:]...)
		grgs := ComputeGRG(ref, valuesTemp)

		grgs = append(grgs, 0)
		copy(grgs[i+1:], grgs[i:])
		grgs[i] = 0

		res[i] = grgs
	}

	return res
}

// ComputeGRG compute single serial's grey relational grade
// @link https://www.cnblogs.com/aabbcc/p/9747715.html
func ComputeGRG(ref []float64, values [][]float64) []float64 {
	r, c := shape(values)
	valuesAbs := copyValuesUseCopy(values)

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			valuesAbs[i][j] = math.Abs(valuesAbs[i][j] - ref[j])
		}
	}

	maxs := make([]float64, r)
	mins := make([]float64, r)
	var max, min float64
	for i := 0; i < r; i++ {
		max, min = maxmin(valuesAbs[i])
		maxs[i] = max
		mins[i] = min
	}

	max, _ = maxmin(maxs)
	_, min = maxmin(mins)

	res := make([]float64, r)

	for i := 0; i < r; i++ {
		var grcSum float64 = 0
		for j := 0; j < c; j++ {
			grcSum = grcSum + (min+Rho*max)/(valuesAbs[i][j]+Rho*max)
		}
		res[i] = 1 / float64(c) * grcSum
	}

	return res
}
