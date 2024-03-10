package strategies

import "prisoners-dilemma/internal/core"

type GrdTFTStrategy struct {
	id core.ID
}

func NewGrdTFTStrategy() *GrdTFTStrategy {
	return &GrdTFTStrategy{
		id: core.GenRandID(),
	}
}

func (s GrdTFTStrategy) ID() core.ID {
	return s.id
}

func (s GrdTFTStrategy) Abbr() string {
	return "GrdTFT"
}

func (s GrdTFTStrategy) Name() string {
	return "Gradual Tit for Tat"
}

func (s GrdTFTStrategy) Description() string {
	return "TFT with two differences: (1) it increases the string of punishing defection responses with each additional defection by its opponent (2) it apologizes for each string of defections by cooperating in the subsequent two rounds"
}

func (s *GrdTFTStrategy) Strategy(state *core.State) (core.Move, error) {
	defectsTotalKey := "defects-total"
	defectsCounterKey := "defects-counter"

	if !state.Self().Store().Has(defectsTotalKey) {
		state.Self().Store().Set(defectsTotalKey, 0)
	}

	if !state.Self().Store().Has(defectsCounterKey) {
		state.Self().Store().Set(defectsCounterKey, 0)
	}

	if state.Iteration() > 0 {
		move := state.FirstOponent().Move(-1)
		if move == core.MOVE_DEFECT {
			total, _ := cast[int](state, defectsTotalKey)
			state.Self().Store().Set(defectsTotalKey, total+1)
			state.Self().Store().Set(defectsCounterKey, total+1)
		}
		return move, nil
	}

	counter, _ := cast[int](state, defectsCounterKey)
	if counter > 0 {
		state.Self().Store().Set(defectsCounterKey, counter-1)
		return core.MOVE_DEFECT, nil
	}
	return core.MOVE_COOPERATE, nil
}

func cast[T any](state *core.State, key string) (T, bool) {
	value, ok := state.Self().Store().Get(key)
	if !ok {
		var zero T
		return zero, false
	}
	val, ok := value.(T)
	if !ok {
		var zero T
		return zero, false
	}
	return val, true
}
