package strategies

import "prisoners-dilemma/internal/core"

type GoodStrategy struct {
	id core.ID
}

func NewGoodStrategy() *GoodStrategy {
	return &GoodStrategy{
		id: core.GenRandID(),
	}
}

func (s GoodStrategy) ID() core.ID {
	return s.id
}

func (s GoodStrategy) Abbr() string {
	return "GOOD"
}

func (s GoodStrategy) Name() string {
	return "Good"
}

func (s GoodStrategy) Description() string {
	return `A Good strategy for the infinitely-repeated, two-player PD is a strategy with the following properties:
(1) its use by both players ensures that each gets reward as long-term average payoff
(2) it is a nash-equilibrium with itself
(3) if it is employed by both, any deviation by one that reduces the average payoff of the other will also reduce its own average payoff.
Aikin, 2013 provides a simple characterization of the memory-one strategies that are GOOD.`
}

func (s GoodStrategy) Strategy(state *core.State) (core.Move, error) {
	return core.MOVE_COOPERATE, nil
}
