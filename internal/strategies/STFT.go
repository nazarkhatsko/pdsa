package strategies

import (
	"errors"
	"prisoners-dilemma/internal/core"
)

type STFTStrategy struct {
	id core.ID
}

func NewSTFTStrategy() *STFTStrategy {
	return &STFTStrategy{
		id: core.GenRandID(),
	}
}

func (s STFTStrategy) ID() core.ID {
	return s.id
}

func (s STFTStrategy) Abbr() string {
	return "STFT"
}

func (s STFTStrategy) Name() string {
	return "Suspicious Tit for Tat"
}

func (s STFTStrategy) Description() string {
	return "Defects on the first round and imitates its opponent's previous move thereafter"
}

func (s *STFTStrategy) Strategy(state *core.State) (core.Move, error) {
	if state.OponentsCount() != 1 {
		return core.MOVE_NULL, errors.New("error: STFTStrategy() incorrect oponents count")
	}

	if state.Iteration() > 0 {
		move := state.FirstOponent().Move(-1)
		return move, nil
	}
	return core.MOVE_DEFECT, nil
}
