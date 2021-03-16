Gray Relational Analysis Implemented within Golang
========================

`gra` is a Golang package for computing grey relational grade.

### Installation

```bash
go get -u github.com/a2htray/gra
```

In your project:

```go
package main

import "github.com/a2htray/gra"

func main() {
    // ...
}
```

### Usages

Compute a set of grey relational grades between a reference slice and a slice collection

```go
package main

import "github.com/a2htray/gra"

func main() {
    values := [][]float64{
		{8, 9, 8, 7, 5, 2, 9},
		{7, 8, 7, 5, 7, 3, 8},
		{9, 7, 9, 6, 6, 4, 7},
		{6, 8, 8, 8, 4, 3, 6},
		{8, 6, 6, 9, 8, 3, 8},
		{8, 9, 5, 7, 6, 4, 8},
	}
	ref := []float64{9, 9, 9, 9, 8, 9, 9}

	fmt.Println(gra.ComputeGRG(ref, values))
}
```

```bash
[0.7233877233877233 0.6344497607655503 0.6941881647764001 0.6064777327935222 0.7144142407300302 0.6723877429759783]
```

Show in table (fixed in `3` precision):

|instance|x0|x1|x2|x3|x4|x5|
|:---|:---|:---|:---|:---|:---|:---|
|ref|0.723|0.634|0.694|0.606|0.714|0.672|

From the above table, we can infer that the relation of which instance is the nearest to the ref instace is `x0 > x4 > x2 > x5 > x1 > x3`.

Compute a matrix of grey relational grades with a given dataset. For example, choosing a subset of features which contribute to the target (`Feature Selection`) is a important step in machine learning problem.

|f0|f1|f2|f3|f4|f5|f6|
|:---|:---|:---|:---|:---|:---|:---|
|8|9|8|7|5|2|9|
|7|8|7|5|7|3|8|
|9|7|9|6|6|4|7|
|6|8|8|8|4|3|6|
|8|6|6|9|8|3|8|
|8|9|5|7|6|4|8|
|9|9|9|9|8|9|9|

```go
package main

import "github.com/a2htray/gra"

func main() {
    values := [][]float64{
		{8, 9, 8, 7, 5, 2, 9},
		{7, 8, 7, 5, 7, 3, 8},
		{9, 7, 9, 6, 6, 4, 7},
		{6, 8, 8, 8, 4, 3, 6},
		{8, 6, 6, 9, 8, 3, 8},
		{8, 9, 5, 7, 6, 4, 8},
		{9, 9, 9, 9, 8, 9, 9},
	}

	grgs := gra.ComputeGRGs(values, gra.AxisColumn)

    for _, grg := range grgs {
		fmt.Println(grg)
	}
}
```

```bash
[0 0.63718820861678 0.7442874585731728 0.5303109588823874 0.6672850958565244 0.6155590441304727 0.776953205524634]
[0.7214285714285715 0 0.7035714285714285 0.669047619047619 0.6476190476190475 0.7107142857142856 0.8142857142857143]
[0.7714285714285715 0.7380952380952379 0 0.6047619047619047 0.6285714285714284 0.6904761904761906 0.6666666666666665]
[0.6285714285714284 0.638095238095238 0.638095238095238 0 0.6666666666666665 0.5428571428571427 0.5809523809523809]
[0.6763553906411048 0.560312703169846 0.6111190396904683 0.6626548055119484 0 0.6128633271490413 0.5566972709829853]
[0.5233921908844509 0.5207929325576384 0.5678366353288953 0.5718770393692992 0.6064536770419123 0 0.4722099366062214]
[0.8845598845598845 0.8643578643578642 0.7147297147297147 0.7349317349317349 0.7246753246753246 0.5286360698125403 0]
```

|feature<br />target feature|f0|f1|f2|f3|f4|f5|f6|
|:---|:---|:---|:---|:---|:---|:---|:---|
|f0|0|0.637|0.744|0.530|0.667|0.615|0.776|
|f1|0.721|0|0.703|0.669|0.647|0.710|0.814|
|f2|0.771|0.738|0|0.604|0.628|0.690|0.666|
|f3|0.628|0.638|0.638|0|0.666|0.542|0.580|
|f4|0.676|0.560|0.611|0.662|0|0.612|0.556|
|f5|0.523|0.520|0.567|0.571|0.606|0|0.472|
|f6|0.884|0.864|0.714|0.734|0.724|0.528|0|

The larger the value, the stronger the relationship with the target feature.