package core

import "prisoners-dilemma/internal/utils"

type StateManager struct {
	iteration int
	ids       []ID
	balances  map[ID]Point
	moves     map[ID][]Move
	stores    map[ID]StateStore
}

func NewStateManager(ids []ID) *StateManager {
	manager := StateManager{
		iteration: 0,
		ids:       ids,
		balances:  make(map[ID]Point),
		moves:     make(map[ID][]Move, 0),
		stores:    make(map[ID]StateStore),
	}

	for _, id := range ids {
		manager.balances[id] = 0
		manager.moves[id] = make([]Move, 0)
		manager.stores[id] = StateStore{
			store: make(map[string]any),
		}
	}

	return &manager
}

func (sm *StateManager) SwitchTo(id ID) *State {
	store := sm.stores[id]

	oponents := utils.Accumulate(sm.ids, func(id ID, acc []StateOponent) []StateOponent {
		return append(acc, StateOponent{
			StateContext: StateContext{
				balance: sm.balances[id],
				moves:   sm.moves[id],
			},
		})
	}, make([]StateOponent, 0))

	state := State{
		iteration: sm.iteration,
		self: StateSelf{
			StateContext: StateContext{
				balance: sm.balances[id],
				moves:   sm.moves[id],
			},
			store: &store,
		},
		oponents: oponents,
	}

	return &state
}

func (sm *StateManager) Update(points []Point, moves []Move) {
	sm.iteration++

	for i, id := range sm.ids {
		sm.balances[id] += points[i]
		sm.moves[id] = append(sm.moves[id], moves[i])
	}
}
