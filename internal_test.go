package gra

import (
	"fmt"
	"testing"
)

func Test_maxmin(t *testing.T) {
	values := []float64{1, 2, 3, 4, 5, 6, 7, -1}
	max, min := maxmin(values)

	if max != 7 || min != -1 {
		t.Fatal("maxmin func got a bad ans")
	}
}

func Test_shape(t *testing.T) {
	values := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}

	if r, c := shape(values); r != 2 || c != 3 {
		t.Fatal("shape is incorrect")
	}
}

func Test_t(_ *testing.T) {
	values := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("before", values)
	fmt.Println("after", t(values))
}

func Test_zeros(t *testing.T) {
	values := zeros(4, 3)
	fmt.Println(values)
}

func Test_copyValuesUseLoop(t *testing.T) {
	values := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}

	valuesCopy := copyValuesUseLoop(values)

	fmt.Println("origin values addr", fmt.Sprintf("%p", &values))
	fmt.Println("origin values", values)

	fmt.Println("copyed values addr", fmt.Sprintf("%p", &valuesCopy))
	fmt.Println("copyed values", valuesCopy)
}

func Test_copyValuesUseCopy(t *testing.T) {
	values := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}

	valuesCopy := copyValuesUseCopy(values)

	fmt.Println("origin values addr", fmt.Sprintf("%p", &values))
	fmt.Println("origin values", values)

	fmt.Println("copyed values addr", fmt.Sprintf("%p", &valuesCopy))
	fmt.Println("copyed values", valuesCopy)
}

func Benchmark_copyValuesUseLoop(b *testing.B) {
	values := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}

	for i := 0; i < 10000; i++ {
		_ = copyValuesUseLoop(values)
	}
}

func Benchmark_copyValuesUseCopy(b *testing.B) {
	values := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}

	for i := 0; i < 10000; i++ {
		_ = copyValuesUseCopy(values)
	}
}
