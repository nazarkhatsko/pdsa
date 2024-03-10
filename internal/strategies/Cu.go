package strategies

import "prisoners-dilemma/internal/core"

type CuStrategy struct {
	id core.ID
}

func NewCuStrategy() *CuStrategy {
	return &CuStrategy{
		id: core.GenRandID(),
	}
}

func (s CuStrategy) ID() core.ID {
	return s.id
}

func (s CuStrategy) Abbr() string {
	return "Cu"
}

func (s CuStrategy) Name() string {
	return "Unconditional Cooperator"
}

func (s CuStrategy) Description() string {
	return "Cooperates unconditionally"
}

func (s CuStrategy) Strategy(state *core.State) (core.Move, error) {
	return core.MOVE_COOPERATE, nil
}
