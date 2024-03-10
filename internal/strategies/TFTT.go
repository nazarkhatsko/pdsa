package strategies

import (
	"errors"
	"prisoners-dilemma/internal/core"
)

type TFTTStrategy struct {
	id core.ID
}

func NewTFTTStrategy() *TFTTStrategy {
	return &TFTTStrategy{
		id: core.GenRandID(),
	}
}

func (s TFTTStrategy) ID() core.ID {
	return s.id
}

func (s TFTTStrategy) Abbr() string {
	return "TFTT"
}

func (s TFTTStrategy) Name() string {
	return "Tit for Two Tats"
}

func (s TFTTStrategy) Description() string {
	return "Cooperates unless defected against twice in a row"
}

func (s *TFTTStrategy) Strategy(state *core.State) (core.Move, error) {
	if state.OponentsCount() != 1 {
		return core.MOVE_NULL, errors.New("error: TFTTStrategy() incorrect oponents count")
	}

	if state.Iteration() > 1 {
		move0 := state.FirstOponent().Move(-1)
		move1 := state.FirstOponent().Move(-2)

		if move0 == core.MOVE_DEFECT && move1 == core.MOVE_DEFECT {
			return core.MOVE_DEFECT, nil
		}
	}
	return core.MOVE_COOPERATE, nil
}
