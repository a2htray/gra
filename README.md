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
|9|9|9|9|8|9|9|
|8|9|8|7|5|2|9|
|7|8|7|5|7|3|8|
|9|7|9|6|6|4|7|
|6|8|8|8|4|3|6|
|8|6|6|9|8|3|8|
|8|9|5|7|6|4|8|

```go
package main

import "github.com/a2htray/gra"

func main() {
    values := [][]float64{
        {9, 9, 9, 9, 8, 9, 9},
        {8, 9, 8, 7, 5, 2, 9},
        {7, 8, 7, 5, 7, 3, 8},
        {9, 7, 9, 6, 6, 4, 7},
        {6, 8, 8, 8, 4, 3, 6},
        {8, 6, 6, 9, 8, 3, 8},
        {8, 9, 5, 7, 6, 4, 8},
    }
    grgs := gra.ComputeGRGs(values, gra.AxisColumn)

    for i := 0; i < len(grgs); i++ {
        fmt.Printf("ref-feature#%d\t", i)
        for j := 0; j < len(grgs[0]); j++ {
            fmt.Printf("%.6f\t", grgs[i][j])
        }
        fmt.Println()
    }

    fmt.Println()

    grgs = gra.ComputeGRGs(t(values), gra.AxisRow)
    for i := 0; i < len(grgs); i++ {
        fmt.Printf("ref-feature#%d\t", i)
        for j := 0; j < len(grgs[0]); j++ {
            fmt.Printf("%.6f\t", grgs[i][j])
        }
        fmt.Println()
    }
}
```

```bash
ref-feature#0   0.000000        0.721429        0.814286        0.707143        0.707143        0.491497        0.871429
ref-feature#1   0.748918        0.000000        0.808369        0.732490        0.634499        0.520793        0.864358
ref-feature#2   0.814286        0.789796        0.000000        0.707143        0.646939        0.537415        0.685714
ref-feature#3   0.707143        0.707143        0.707143        0.000000        0.696939        0.540476        0.707143
ref-feature#4   0.676355        0.560313        0.611119        0.662655        0.000000        0.530311        0.662655
ref-feature#5   0.523392        0.520793        0.567837        0.571877        0.606454        0.000000        0.528636
ref-feature#6   0.884560        0.864358        0.714730        0.734932        0.724675        0.528636        0.000000

ref-feature#0   0.000000        0.721429        0.814286        0.707143        0.707143        0.491497        0.871429
ref-feature#1   0.748918        0.000000        0.808369        0.732490        0.634499        0.520793        0.864358
ref-feature#2   0.814286        0.789796        0.000000        0.707143        0.646939        0.537415        0.685714
ref-feature#3   0.707143        0.707143        0.707143        0.000000        0.696939        0.540476        0.707143
ref-feature#4   0.676355        0.560313        0.611119        0.662655        0.000000        0.530311        0.662655
ref-feature#5   0.523392        0.520793        0.567837        0.571877        0.606454        0.000000        0.528636
ref-feature#6   0.884560        0.864358        0.714730        0.734932        0.724675        0.528636        0.000000
```

The larger the value, the stronger the relationship with the target feature.

### Rho

Using multiple `Rho`s:

```go
package main

import "github.com/a2htray/gra"

func main() {
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
        gra.Rho = rho
        fmt.Println("rho:", rho)
        fmt.Println(gra.ComputeGRG(ref, values))
    }
}
```

```bash
rho: 0.1
[0.4804124215888921 0.28674663118837734 0.4413962835015466 0.2667266111683574 0.4723407717261975 0.37339873114854993]
rho: 0.3
[0.6442239668971745 0.5228814089656849 0.6063035313747621 0.49418700437458024 0.6339467686716264 0.5741808423641642]
rho: 0.5
[0.7233877233877233 0.6344497607655503 0.6941881647764001 0.6064777327935222 0.7144142407300302 0.6723877429759783]
rho: 0.7
[0.7718828210182985 0.7017026229939651 0.7493767774450216 0.6760192621427451 0.7644377946671621 0.7324025852753301]
```