package solvers

import "prisoners-dilemma/internal/core"

var Solvers []core.Solver

func init() {
	Solvers = []core.Solver{
		NewSimpleSolver(),
	}
}
