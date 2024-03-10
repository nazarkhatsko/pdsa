package solvers

import (
	"errors"
	"prisoners-dilemma/internal/core"
	"prisoners-dilemma/internal/utils"
)

type SMPSolver struct {
	id core.ID
}

func NewSMPSolver() *SMPSolver {
	return &SMPSolver{
		id: core.GenRandID(),
	}
}

func (s SMPSolver) ID() core.ID {
	return s.id
}

func (s SMPSolver) Abbr() string {
	return "SMP"
}

func (s SMPSolver) Name() string {
	return "Simple"
}

func (s SMPSolver) Description() string {
	return "Classic Solver for resolving disputes between 2 players"
}

func (s SMPSolver) Solver(moves []core.Move) ([]core.Point, error) {
	if len(moves) != 2 {
		return utils.List[core.Point](), errors.New("error: SMPSolver() incorrect moves count")
	}

	if moves[0] == core.MOVE_COOPERATE && moves[1] == core.MOVE_COOPERATE {
		return utils.List[core.Point](3, 3), nil
	} else if moves[0] == core.MOVE_COOPERATE && moves[1] == core.MOVE_DEFECT {
		return utils.List[core.Point](0, 5), nil
	} else if moves[0] == core.MOVE_DEFECT && moves[1] == core.MOVE_COOPERATE {
		return utils.List[core.Point](5, 0), nil
	} else if moves[0] == core.MOVE_DEFECT && moves[1] == core.MOVE_DEFECT {
		return utils.List[core.Point](1, 1), nil
	}

	return utils.List[core.Point](), errors.New("error: SMPSolver() no solution found")
}
