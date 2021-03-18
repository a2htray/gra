package gra

import (
	"fmt"
	"testing"
)

func Test_ComputeGRG(t *testing.T) {
	// compute the grades between the ref sample and each of sample in dataset
	// the similiarity metric is grey relational grade
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
	// compute the similiarities in all samples
	// one row represents a sample
	values := [][]float64{
		{9, 9, 9, 9, 8, 9, 9},
		{8, 9, 8, 7, 5, 2, 9},
		{7, 8, 7, 5, 7, 3, 8},
		{9, 7, 9, 6, 6, 4, 7},
		{6, 8, 8, 8, 4, 3, 6},
		{8, 6, 6, 9, 8, 3, 8},
		{8, 9, 5, 7, 6, 4, 8},
	}

	grgs := ComputeGRGs(values, AxisRow)
	for _, grg := range grgs {
		fmt.Println(grg)
	}
}

func Test_ComputeGRG2(_t *testing.T) {
	// compute each feature's contribution to the target feature
	values := [][]float64{
		{9, 9, 9, 9, 8, 9},
		{8, 9, 8, 7, 5, 2},
		{7, 8, 7, 5, 7, 3},
		{9, 7, 9, 6, 6, 4},
		{6, 8, 8, 8, 4, 3},
		{8, 6, 6, 9, 8, 3},
		{8, 9, 5, 7, 6, 4},
	}
	ref := []float64{9, 9, 8, 7, 6, 8, 8}

	fmt.Println(ComputeGRG(ref, t(values)))
}

func Test_ComputeGRGs2(_t *testing.T) {
	// compute all contributions within all features
	values := [][]float64{
		{9, 9, 9, 9, 8, 9, 9},
		{8, 9, 8, 7, 5, 2, 9},
		{7, 8, 7, 5, 7, 3, 8},
		{9, 7, 9, 6, 6, 4, 7},
		{6, 8, 8, 8, 4, 3, 6},
		{8, 6, 6, 9, 8, 3, 8},
		{8, 9, 5, 7, 6, 4, 8},
	}
	grgs := ComputeGRGs(values, AxisColumn)

	for i := 0; i < len(grgs); i++ {
		fmt.Printf("ref-feature#%d\t", i)
		for j := 0; j < len(grgs[0]); j++ {
			fmt.Printf("%.6f\t", grgs[i][j])
		}
		fmt.Println()
	}

	fmt.Println()

	grgs = ComputeGRGs(t(values), AxisRow)
	for i := 0; i < len(grgs); i++ {
		fmt.Printf("ref-feature#%d\t", i)
		for j := 0; j < len(grgs[0]); j++ {
			fmt.Printf("%.6f\t", grgs[i][j])
		}
		fmt.Println()
	}
}

func Test_UseMultipleRho(t *testing.T) {
	rhos := []float64{0.1, 0.3, 0.5, 0.7}

	values := [][]float64{
		{8, 9, 8, 7, 5, 2, 9},
		{7, 8, 7, 5, 7, 3, 8},
		{9, 7, 9, 6, 6, 4, 7},
		{6, 8, 8, 8, 4, 3, 6},
		{8, 6, 6, 9, 8, 3, 8},
		{8, 9, 5, 7, 6, 4, 8},
	}
	ref := []float64{9, 9, 9, 9, 8, 9, 9}

	for _, rho := range rhos {
		Rho = rho
		fmt.Println("rho:", rho)
		fmt.Println(ComputeGRG(ref, values))
	}
}
