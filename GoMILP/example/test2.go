package main

import (
	"context"
	"fmt"
	ilp "github.com/baconpb/algo/GoMILP"
	"gonum.org/v1/gonum/mat"
	"time"
)

type NewMilp struct {
	my_obj []float64
	my_colnames []string
	rows_eq [][]float64
	cons_eq_ret []float64
	my_rhs_eq []float64
	rows_ineq [][]float64
	my_rhs_ineq []float64
	cons_ineq_ret []float64
	my_ub []float64
	my_lb []float64
	my_ctype []bool
}

func Test()  {
	var newMilps []NewMilp
	var newMilp NewMilp
	newMilp.rows_eq = [][]float64{
		{3, 1, 0, 1},
	}
	newMilp.cons_eq_ret = []float64{3, 1, 0, 1}
	newMilp.my_rhs_eq = []float64{9.2}
	newMilp.rows_ineq = [][]float64{
		{-0.5, 0, 1.5, 0},
		{3.1, 5, 0, 1.1},
		{1.1, 0, 0.5, 2.1},
		{-2.1, -1.2, 0, -1.1},
	}
	newMilp.cons_ineq_ret = []float64{
		-0.5, 0, 1.5, 0,
		3.1, 5, 0, 1.1,
		1.1, 0, 0.5, 2.1,
		-2.1, -1.2, 0, -1.1,
	}
	newMilp.my_rhs_ineq = []float64{2.8,25.5, 20.6,20,30}
	newMilps = append(newMilps, newMilp)
	TestMilpAuto(newMilps)
}

func TestMilpAuto(newMilps []NewMilp) {
	my_obj := []float64{-1.2,2.3,5.2,-1.5}
	my_colnames := []string{"x1","x2","x3","x4"}
	rows_eq := [][]float64{
		{-1, 2, 1, 0},
	}

	cons_eq_ret := []float64{
		-1, 2, 1, 0,
	}
	my_rhs_eq := []float64{4}
	rows_ineq := [][]float64{
		{-1,0,0,0},
		{1,0,0,0},
		{0,-1,0,0},
		{0,1,0,0},
		{0,0,-1,0},
		{0,0,1,0},
		{0,0,0,-1},
		{0,0,0,1},

	}
	my_rhs_ineq := []float64{12,12.5,10,10.8,10,10.5,2}
	cons_ineq_ret :=[]float64{
		-1,0,0,0,
		1,0,0,0,
		0,-1,0,0,
		0,1,0,0,
		0,0,-1,0,
		0,0,1,0,
		0,0,0,-1,
		0,0,0,1,

	}
	//my_ub := []float64{}
	//my_lb := []float64{}
	my_ctype := []bool{false,false,false,false}

	//配置
	for _,oneMilp := range newMilps {
		for _,item := range oneMilp.rows_eq {
			rows_eq = append(rows_eq, item)
		}
		for _,item := range oneMilp.cons_eq_ret {
			cons_eq_ret = append(cons_eq_ret, item)
		}
		for _,item := range oneMilp.my_rhs_eq {
			my_rhs_eq = append(my_rhs_eq, item)
		}
		for _,item := range oneMilp.rows_ineq {
			rows_ineq = append(rows_ineq, item)
		}
		for _,item := range oneMilp.my_rhs_ineq {
			my_rhs_ineq = append(my_rhs_ineq, item)
		}
		for _,item := range oneMilp.cons_ineq_ret {
			cons_ineq_ret = append(cons_ineq_ret, item)
		}
	}

	prob := ilp.MilpProblem{
		C: my_obj, //目标函数系数  min(c1*q1+c2*e1+c2*e3+c3*e5-c4*e2-c4*e4)
		A: mat.NewDense(len(rows_eq), len(my_colnames), cons_eq_ret),
		B: my_rhs_eq,
		//G: mat.NewDense(len(rows_ineq)+len(my_ub)*2, len(my_colnames), cons_ineq_ret),
		G: mat.NewDense(len(rows_ineq), len(my_colnames), cons_ineq_ret),
		H: my_rhs_ineq,
		//h: []float64{25.5, 20.6,20,30},
		IntegralityConstraints: my_ctype,
	}

	// Solve the problem with 1 worker and a one-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	got, _ := prob.Solve(ctx, 1, ilp.DummyMiddleware{})
	fmt.Println(prob.G)
	fmt.Println(got.Z)
	fmt.Println(got.X)
	// dump the logged tree to a DOT-file
}

func TestMilpAuto2() {
	//prob := MilpProblem{
	//	c: []float64{-1.2,2.3,5.2,-1.5}, //目标函数系数  min(c1*q1+c2*e1+c2*e3+c3*e5-c4*e2-c4*e4)
	//	A: mat.NewDense(2, 4, []float64{
	//		-1, 2, 1, 0,
	//		3, 1, 0, 1,
	//	}),
	//	b: []float64{4, 9.2},
	//	G: mat.NewDense(12, 4, []float64{
	//		-1,0,0,0,
	//		1,0,0,0,
	//		0,-1,0,0,
	//		0,1,0,0,
	//		0,0,-1,0,
	//		0,0,1,0,
	//		0,0,0,-1,
	//		0,0,0,1,
	//		-0.5, 0, 1.5, 0,
	//		3.1, 5, 0, 1.1,
	//		1.1, 0, 0.5, 2.1,
	//		-2.1, -1.2, 0, -1.1,
	//	}),
	//	h: []float64{12,12.5,10,10.8,10,10.5,2,2.8,25.5, 20.6,20,30},
	//	//h: []float64{25.5, 20.6,20,30},
	//	IntegralityConstraints: []bool{false, false, false, false},
	//}
	//my_ctype_num := []int{18,2}
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
		//G: mat.NewDense(len(rows_ineq)+len(my_ub)*2, len(my_colnames), cons_ineq_ret),
		G: mat.NewDense(len(rows_ineq), len(my_colnames), cons_ineq_ret),
		H: my_rhs_ineq,
		//h: []float64{25.5, 20.6,20,30},
		IntegralityConstraints: my_ctype,
	}

	// Solve the problem with 1 worker and a one-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	got, _ := prob.Solve(ctx, 1, ilp.DummyMiddleware{})
	fmt.Println(prob.G)
	fmt.Println(got.Z)
	fmt.Println(got.X)
	// dump the logged tree to a DOT-file
}

func main()  {
	Test()
	TestMilpAuto2()
}