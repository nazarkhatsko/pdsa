package solvers

import (
	"errors"
	"prisoners-dilemma/internal/core"
	"prisoners-dilemma/internal/utils"
)

type SimpleSolver struct {
	id core.ID
}

func NewSimpleSolver() *SimpleSolver {
	return &SimpleSolver{
		id: core.GenRandID(),
	}
}

func (s SimpleSolver) ID() core.ID {
	return s.id
}

func (s SimpleSolver) Abbr() string {
	return "SMP"
}

func (s SimpleSolver) Name() string {
	return "Simple"
}

func (s SimpleSolver) Description() string {
	return "Description"
}

func (s SimpleSolver) Solver(moves []core.Move) ([]core.Point, error) {
	if len(moves) != 2 {
		return utils.List[core.Point](), errors.New("error: SimpleSolver() incorrect moves")
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

	return utils.List[core.Point](), errors.New("error: SimpleSolver()")
}
