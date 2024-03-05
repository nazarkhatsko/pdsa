package strategies

import "prisoners-dilemma/internal/core"

type TFTStrategy struct {
	id core.ID
}

func NewTFTStrategy() *TFTStrategy {
	return &TFTStrategy{
		id: core.GenRandID(),
	}
}

func (s TFTStrategy) ID() core.ID {
	return s.id
}

func (s TFTStrategy) Abbr() string {
	return "TFT"
}

func (s TFTStrategy) Name() string {
	return "Tit for Tat"
}

func (s TFTStrategy) Description() string {
	return "Cooperates on the first round and imitates its opponent's previous move thereafter."
}

func (s *TFTStrategy) Strategy(state *core.State) (core.Move, error) {
	if state.Iteration() > 0 {
		move, _ := state.FirstOponent().LastMove()
		return move, nil
	}
	return core.MOVE_COOPERATE, nil
}
