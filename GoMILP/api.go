package ilp

import (
	"context"
	"math"

	"gonum.org/v1/gonum/mat"
)

// The abstract MILP problem representation
type Problem struct {
	// minimizes by default
	maximize bool

	// the problem structure
	variables   []*Variable
	constraints []*Constraint

	// the branching heuristic to use for branch-and-bound (defaults to 0 == maxFun)
	branchingHeuristic BranchHeuristic

	// number of workers to Solve the MilpProblem with
	workers int

	// instrumentation middleware
	instrumentation BnbMiddleware
}

// A variable of the MILP problem.
type Variable struct {
	// variable name for human reference
	name string

	// coefficient of the variable in the objective function
	coefficient float64

	// integrality constraint
	integer bool

	// bounds
	upper float64
	lower float64
}

// an expression of a variable and an arbitrary float for use in defining constraints
// e.g. "-1 * x1"
type expression struct {
	coef     float64
	variable *Variable
}

type Constraint struct {
	// these expressions will be summed together to form the left-hand-side of the constraint
	expressions []expression

	// right-hand side of the constraint equation
	rhs float64

	// an equality constraint by default
	inequality bool

	// store a reference to the problem
	problem *Problem
}

// Initiate a new MILP problem abstraction
func NewProblem() Problem {
	return Problem{
		workers:         1,
		instrumentation: DummyMiddleware{},
	}
}

// add a variable and return a reference to that variable.
// Defaults to no integrality constraint and an objective function coefficient of 0
func (p *Problem) AddVariable(name string) *Variable {
	// TODO: check for uniqueness of variable name as it has an important role downstream

	v := Variable{
		name:        name,
		coefficient: 0,
		integer:     false,
		upper:       math.Inf(1),
		lower:       0,
	}

	p.variables = append(p.variables, &v)

	return &v
}

// SetCoeff sets the value of the variable in the objective function
func (v *Variable) SetCoeff(coef float64) *Variable {
	v.coefficient = coef
	return v
}

func (v *Variable) IsInteger() *Variable {
	v.integer = true
	return v
}

// UpperBound sets the inclusive upper bound of this variable. Input must be positive.
func (v *Variable) UpperBound(bound float64) *Variable {
	v.upper = bound
	return v
}

// LowerBound sets the inclusive lower bound of this variable. Input must be positive.
func (v *Variable) LowerBound(bound float64) *Variable {
	v.lower = bound
	return v
}

func (p *Problem) AddConstraint() *Constraint {
	c := &Constraint{
		problem: p,
	}
	p.constraints = append(p.constraints, c)

	return c
}

func (p *Constraint) EqualTo(val float64) *Constraint {
	p.inequality = false
	p.rhs = val
	return p
}

func (p *Constraint) SmallerThanOrEqualTo(val float64) *Constraint {
	p.inequality = true
	p.rhs = val
	return p
}

func (c *Constraint) AddExpression(coef float64, v *Variable) *Constraint {
	// check if the provided variable has been declared in this problem. If not, this call will panic
	c.problem.getVariableIndex(v)

	exp := expression{coef: coef, variable: v}

	c.expressions = append(c.expressions, exp)
	return c
}

func (p *Problem) Maximize() {
	p.maximize = true
}

func (p *Problem) Minimize() {
	p.maximize = false
}

func (p *Problem) BranchingHeuristic(choice BranchHeuristic) {
	p.branchingHeuristic = choice
}

func (p *Problem) SetWorkers(n int) {
	p.workers = n
}

func (p *Problem) SetInstrumentation(b BnbMiddleware) {
	p.instrumentation = b
}

// Check whether the expression is legal considering the variables currently present in the problem
func (p *Problem) checkExpression(e expression) bool {

	// check whether the pointer to the variable provided is currently included in the Problem
	for _, v := range p.variables {
		if v == e.variable {
			return true
		}
	}

	return false

}

// get the index of the variable pointer in the variable pointer slice of the Problem struct using a linear search
func (p *Problem) getVariableIndex(v *Variable) int {
	for i, va := range p.variables {
		if v == va {
			return i
		}
	}
	panic("variable pointer not found in Problem struct")
}

// Convert the abstract problem representation to its concrete numerical representation.
func (p Problem) toSolveable() *MilpProblem {

	// get the c vector containing the coefficients of the variables in the objective function
	// simultaneously parse the integrality constraints
	var c []float64
	var integrality []bool
	for _, v := range p.variables {

		// if the Problem is set to be maximized, we assume that all variable coefficients reflect that.
		// To turn this maximization problem into a minimization one, we multiply all coefficients with -1.
		k := v.coefficient
		if p.maximize {
			k = k * -1
		}

		c = append(c, k)
		integrality = append(integrality, v.integer)
	}

	/// parse the constraints
	var b []float64
	var Adata []float64
	var h []float64
	var Gdata []float64
	for _, constraint := range p.constraints {

		// build the matrix row
		indexRow := make([]float64, len(p.variables))

		for _, exp := range constraint.expressions {
			i := p.getVariableIndex(exp.variable)
			indexRow[i] = exp.coef
		}

		if constraint.inequality {
			Gdata = append(Gdata, indexRow...)

			// add the RHS of the inequality to the h vector
			h = append(h, constraint.rhs)
		} else {
			Adata = append(Adata, indexRow...)
			// add the RHS of the equality to the b vector
			b = append(b, constraint.rhs)
		}

	}

	// combine the Adata vector into a matrix
	var A *mat.Dense
	if len(b) > 0 {
		A = mat.NewDense(len(b), len(p.variables), Adata)
	}

	// add the variable bounds as inequality constraints
	for _, v := range p.variables {

		// convert the upper bound to a row in the constraint matrix
		if !math.IsInf(v.upper, 1) {
			uRow := make([]float64, len(p.variables))
			i := p.getVariableIndex(v)
			uRow[i] = 1

			Gdata = append(Gdata, uRow...)

			// add the RHS of the inequality to the h vector
			h = append(h, v.upper)
		}

		// convert the lower bound to a row in the constraint matrix
		// but ONLY if it is nonzero and nonnegative, because negative domain is illegal and the lower bound of the variable is zero by default.
		if !(v.lower <= 0) {
			uRow := make([]float64, len(p.variables))
			i := p.getVariableIndex(v)
			uRow[i] = -1

			Gdata = append(Gdata, uRow...)

			// add the RHS of the inequality to the h vector
			h = append(h, -v.lower)
		}

	}

	// combine the Gdata vector into a matrix
	var G *mat.Dense
	if len(h) > 0 {
		G = mat.NewDense(len(h), len(p.variables), Gdata)
	}

	return &MilpProblem{
		C: c,
		A: A,
		B: b,
		G: G,
		H: h,
		IntegralityConstraints: integrality,
		branchingHeuristic:     p.branchingHeuristic,
	}
}

// SolveWithCtx converts the abstract Problem to a MilpProblem, Solves it, and parses its output.
// Context requires a context.Context as an argument to govern cancellation and Solve deadlines.
func (p Problem) SolveWithCtx(ctx context.Context) (*Solution, error) {

	preprocessor := newPreprocessor()
	prepped := preprocessor.preSolve(p)

	milp := prepped.toSolveable()

	subSolution, err := milp.Solve(ctx, prepped.workers, prepped.instrumentation)
	if err != nil {
		return nil, err
	}

	// convert the solution vector to a rawSolution by mapping each solution coefficient to the corresponding variable name
	rawSol := make(map[string]float64)
	for i, v := range prepped.variables {
		rawSol[v.name] = subSolution.X[i]
	}

	// postprocess the solution
	soln := preprocessor.postSolve(rawSol)

	return &soln, nil

}

// Solve converts the abstract Problem to a MilpProblem, Solves it, and parses its output.
// It does not time out.
func (p *Problem) Solve() (*Solution, error) {

	return p.SolveWithCtx(context.Background())

}
