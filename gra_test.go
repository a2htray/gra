package gra

import (
	"fmt"
	"testing"
)

func Test_ComputeGRG(t *testing.T) {
	values := [][]float64{
		{8, 9, 8, 7, 5, 2, 9},
		{7, 8, 7, 5, 7, 3, 8},
		{9, 7, 9, 6, 6, 4, 7},
		{6, 8, 8, 8, 4, 3, 6},
		{8, 6, 6, 9, 8, 3, 8},
		{8, 9, 5, 7, 6, 4, 8},
	}
	ref := []float64{9, 9, 9, 9, 8, 9, 9}

	fmt.Println(ComputeGRG(ref, values))
}

func Test_ComputeGRGs(t *testing.T) {
	values := [][]float64{
		{8, 9, 8, 7, 5, 2, 9},
		{7, 8, 7, 5, 7, 3, 8},
		{9, 7, 9, 6, 6, 4, 7},
		{6, 8, 8, 8, 4, 3, 6},
		{8, 6, 6, 9, 8, 3, 8},
		{8, 9, 5, 7, 6, 4, 8},
		{9, 9, 9, 9, 8, 9, 9},
	}

	grgs := ComputeGRGs(values, AxisColumn)
	for _, grg := range grgs {
		fmt.Println(grg)
	}
}

func Test_ComputeGRG2(t *testing.T) {
	values := [][]float64{
		{8, 9, 8, 7, 5, 2, 9},
		{7, 8, 7, 5, 7, 3, 8},
		{9, 7, 9, 6, 6, 4, 7},
		{6, 8, 8, 8, 4, 3, 6},
		{8, 6, 6, 9, 8, 3, 8},
		{8, 9, 5, 7, 6, 4, 8},
	}
	ref := []float64{9, 9, 9, 9, 8, 9, 9}

	rhos := []float64{0.1, 0.3, 0.5, 0.7, 0.9}

	for _, rho := range rhos {
		Rho = rho
		fmt.Println("rho", rho)
		fmt.Println(ComputeGRG(ref, values))
	}
}
