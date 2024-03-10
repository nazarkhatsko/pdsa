package strategies

import "prisoners-dilemma/internal/core"

type DuStrategy struct {
	id core.ID
}

func NewDuStrategy() *DuStrategy {
	return &DuStrategy{
		id: core.GenRandID(),
	}
}

func (s DuStrategy) ID() core.ID {
	return s.id
}

func (s DuStrategy) Abbr() string {
	return "Du"
}

func (s DuStrategy) Name() string {
	return "Unconditional Defector"
}

func (s DuStrategy) Description() string {
	return "Defects unconditionally"
}

func (s DuStrategy) Strategy(state *core.State) (core.Move, error) {
	return core.MOVE_DEFECT, nil
}
