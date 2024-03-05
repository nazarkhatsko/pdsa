package core

type Info interface {
	ID() ID
	Abbr() string
	Name() string
	Description() string
}
