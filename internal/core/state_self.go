package core

type StateSelf struct {
	StateContext
	store *StateStore
}

func (ss *StateSelf) Store() *StateStore {
	return ss.store
}
