package strategies

import (
	"math/rand"
	"prisoners-dilemma/internal/core"
)

type RANDStrategy struct {
	id core.ID
}

func NewRANDStrategy() *RANDStrategy {
	return &RANDStrategy{
		id: core.GenRandID(),
	}
}

func (s RANDStrategy) ID() core.ID {
	return s.id
}

func (s RANDStrategy) Abbr() string {
	return "RAND"
}

func (s RANDStrategy) Name() string {
	return "Random"
}

func (s RANDStrategy) Description() string {
	return "Defects unconditionally"
}

func (s RANDStrategy) Strategy(state *core.State) (core.Move, error) {
	if rand.Intn(100) > 50 {
		return core.MOVE_DEFECT, nil
	}
	return core.MOVE_COOPERATE, nil
}
