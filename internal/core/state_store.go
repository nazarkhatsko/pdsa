package core

type StateStore struct {
	store map[string]any
}

func NewStateStore() *StateStore {
	store := StateStore{
		store: make(map[string]any),
	}

	return &store
}

func (ss StateStore) Keys() []string {
	keys := make([]string, 0, len(ss.store))
	for k := range ss.store {
		keys = append(keys, k)
	}
	return keys
}

func (ss StateStore) Values() []any {
	values := make([]any, 0, len(ss.store))
	for _, v := range ss.store {
		values = append(values, v)
	}
	return values
}

func (ss StateStore) Has(key string) bool {
	_, ok := ss.store[key]
	return ok
}

func (ss StateStore) Get(key string) (any, bool) {
	value, ok := ss.store[key]
	return value, ok
}

func (ss *StateStore) Set(key string, value any) {
	ss.store[key] = value
}

func (ss *StateStore) Delete(key string) bool {
	_, ok := ss.store[key]
	delete(ss.store, key)
	return ok
}

func (ss *StateStore) Clear() {
	for k := range ss.store {
		delete(ss.store, k)
	}
}
