package strategies

import (
	"errors"
	"prisoners-dilemma/internal/core"
)

type TTFTStrategy struct {
	id core.ID
}

func NewTTFTStrategy() *TTFTStrategy {
	return &TTFTStrategy{
		id: core.GenRandID(),
	}
}

func (s TTFTStrategy) ID() core.ID {
	return s.id
}

func (s TTFTStrategy) Abbr() string {
	return "TTFT"
}

func (s TTFTStrategy) Name() string {
	return "Two Tits for Tat"
}

func (s TTFTStrategy) Description() string {
	return "Defects twice after being defected against, otherwise cooperates"
}

func (s *TTFTStrategy) Strategy(state *core.State) (core.Move, error) {
	if state.OponentsCount() != 1 {
		return core.MOVE_NULL, errors.New("error: TTFTStrategy() incorrect oponents count")
	}

	if state.Iteration() > 1 {
		move := state.FirstOponent().Move(-1)

		if move == core.MOVE_DEFECT {
			return core.MOVE_DEFECT, nil
		} else if state.Iteration() > 2 {
			move = state.FirstOponent().Move(-2)
			if move == core.MOVE_DEFECT {
				return core.MOVE_DEFECT, nil
			}
		}
	}
	return core.MOVE_COOPERATE, nil
}
