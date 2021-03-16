package gra

// maxmin get the maxinum and mininum of a float64 slice
func maxmin(values []float64) (max, min float64) {
	min = values[0]
	max = values[0]

	for _, v := range values[1:] {
		if min > v {
			min = v
		}

		if max < v {
			max = v
		}
	}

	return
}

// shape return a shape of a given martix
func shape(values [][]float64) (int, int) {
	return len(values), len(values[0])
}

// zeros return a matrix filled with zeros corresponding to a given shape
func zeros(rowNum, columnNum int) [][]float64 {
	res := make([][]float64, rowNum)
	for i := 0; i < rowNum; i++ {
		res[i] = make([]float64, columnNum)
	}

	return res
}

// t transposition
func t(values [][]float64) [][]float64 {
	r, c := shape(values)
	res := zeros(c, r)

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			res[j][i] = values[i][j]
		}
	}

	return res
}

// copyValuesUseLoop copy values by using loop statement
func copyValuesUseLoop(values [][]float64) [][]float64 {
	r, c := shape(values)

	res := zeros(r, c)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			res[i][j] = values[i][j]
		}
	}

	return res
}

// copyValuesUseCopy copy values by using copy function
func copyValuesUseCopy(values [][]float64) [][]float64 {
	r, c := shape(values)
	res := zeros(r, c)
	for i := 0; i < r; i++ {
		copy(res[i], values[i])
	}
	return res
}
