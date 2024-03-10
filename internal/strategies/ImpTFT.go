package strategies

import (
	"errors"
	"math/rand"
	"prisoners-dilemma/internal/core"
)

type ImpTFTStrategy struct {
	id core.ID
}

func NewImpTFTStrategy() *ImpTFTStrategy {
	return &ImpTFTStrategy{
		id: core.GenRandID(),
	}
}

func (s ImpTFTStrategy) ID() core.ID {
	return s.id
}

func (s ImpTFTStrategy) Abbr() string {
	return "ImpTFT"
}

func (s ImpTFTStrategy) Name() string {
	return "Imperfect Tit for Tat"
}

func (s ImpTFTStrategy) Description() string {
	return "Imitates opponent's last move with high (but less than one) probability"
}

func (s *ImpTFTStrategy) Strategy(state *core.State) (core.Move, error) {
	if state.OponentsCount() != 1 {
		return core.MOVE_NULL, errors.New("error: ImpTFTStrategy() incorrect oponents count")
	}

	if state.Iteration() > 0 {
		if rand.Intn(100) > 50 {
			move := state.FirstOponent().Move(-1)
			return move, nil
		}
	}
	return core.MOVE_DEFECT, nil
}
