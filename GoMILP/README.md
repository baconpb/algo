# Get Started
### 1 install
* go get github.com/baconpb/algo/GoMILP

### 2 demo
```go
package main

import (
	"context"
	"fmt"
	ilp "github.com/baconpb/algo/GoMILP"
	"gonum.org/v1/gonum/mat"
	"time"
)
func Test() {
	my_obj := []float64{-1.2,2.3,5.2,-1.5}
	my_colnames := []string{"x1","x2","x3","x4"}
	rows_eq := [][]float64{
		{-1, 2, 1, 0},
		{3, 1, 0, 1},
	}

	cons_eq_ret := []float64{
		-1, 2, 1, 0,
		3, 1, 0, 1,
	}
	my_rhs_eq := []float64{4, 9.2}
	rows_ineq := [][]float64{
		{-1,0,0,0},
		{1,0,0,0},
		{0,-1,0,0},
		{0,1,0,0},
		{0,0,-1,0},
		{0,0,1,0},
		{0,0,0,-1},
		{0,0,0,1},
		{-0.5, 0, 1.5, 0},
		{3.1, 5, 0, 1.1},
		{1.1, 0, 0.5, 2.1},
		{-2.1, -1.2, 0, -1.1},
	}
	my_rhs_ineq := []float64{12,12.5,10,10.8,10,10.5,2,2.8,25.5, 20.6,20,30}
	cons_ineq_ret :=[]float64{
		-1,0,0,0,
		1,0,0,0,
		0,-1,0,0,
		0,1,0,0,
		0,0,-1,0,
		0,0,1,0,
		0,0,0,-1,
		0,0,0,1,
		-0.5, 0, 1.5, 0,
		3.1, 5, 0, 1.1,
		1.1, 0, 0.5, 2.1,
		-2.1, -1.2, 0, -1.1,
	}
	//my_ub := []float64{}
	//my_lb := []float64{}
	my_ctype := []bool{false,false,false,false}
	prob := ilp.MilpProblem{
		C: my_obj, //目标函数系数  min(c1*q1+c2*e1+c2*e3+c3*e5-c4*e2-c4*e4)
		A: mat.NewDense(len(rows_eq), len(my_colnames), cons_eq_ret),
		B: my_rhs_eq,
		G: mat.NewDense(len(rows_ineq), len(my_colnames), cons_ineq_ret),
		H: my_rhs_ineq,
		IntegralityConstraints: my_ctype,
	}

	// Solve the problem with 1 worker and a one-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	got, _ := prob.Solve(ctx, 1, ilp.DummyMiddleware{})
	fmt.Println(prob.G)
	fmt.Println(got.Z)
	fmt.Println(got.X)
}

func main(){
    Test()
}




```
