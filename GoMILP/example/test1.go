//File  : test1.go.go
//Author: 裴彬
//Date  : 2020/12/4

package main

import (
	//"context"
	"fmt"
	ilp "github.com/baconpb/algo/GoMILP"
	"time"

	//"reflect"
	//"time"
)

//type MilpProblem struct {
//	// 	minimize c^T * x
//	// s.t      G * x <= h
//	//          A * x = b
//	c []float64
//	A *mat.Dense
//	b []float64
//	G *mat.Dense
//	h []float64
//
//	// which variables to apply the integrality constraint to. Should have same order as c.
//	IntegralityConstraints []bool
//
//	// which branching heuristic to use. Determines which integer variable is branched on at each split.
//	// defaults to 0 == maxFun
//	branchingHeuristic ilp.BranchHeuristic
//}

func main () {
	//fmt.Println("hello ")
	//prob := MilpProblem{
	//	c: []float64{1.7356332566545616, -0.2058339272568599, -1.051665297603944},
	//	A: mat.NewDense(1, 3, []float64{
	//		-0.7762132098737671, 1.42027949678888, -0.3304567624749696,
	//	}),
	//	b: []float64{-0.24703471683023603},
	//	G: mat.NewDense(1, 3, []float64{
	//		-0.6775235462631393, -1.9616379110849085, 1.9859192819811322,
	//	}),
	//	h: []float64{-0.041138108068992485},
	//	IntegralityConstraints: []bool{true, true, true},
	//}
	//
	////want := ilp.Solution{}
	//
	//// initiate the logger instrumentation
	//tl := ilp.NewTreeLogger()
	//
	//// Solve the problem with 2 workers and a one-second timeout
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//got, _ := prob.Solve(ctx, 2, tl)
	//fmt.Println(got)

	//线性规划案例

	// 混合整数线性规划案例
	t1 := time.Now()
	//ilp.Test_ilp()
	ilp.Test_nilp()
	fmt.Println(10*time.Since(t1))


	ilp.TestMilpAuto2()
	ilp.Test()


}