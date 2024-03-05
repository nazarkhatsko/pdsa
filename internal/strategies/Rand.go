package strategies

import (
	"math/rand"
	"prisoners-dilemma/internal/core"
)

type RandStrategy struct {
	id core.ID
}

func NewRandStrategy() *RandStrategy {
	return &RandStrategy{
		id: core.GenRandID(),
	}
}

func (s RandStrategy) ID() core.ID {
	return s.id
}

func (s RandStrategy) Abbr() string {
	return "Rand"
}

func (s RandStrategy) Name() string {
	return "Random"
}

func (s RandStrategy) Description() string {
	return "Description"
}

func (s RandStrategy) Strategy(state *core.State) (core.Move, error) {
	if rand.Intn(100) > 50 {
		return core.MOVE_COOPERATE, nil
	}
	return core.MOVE_DEFECT, nil
}
