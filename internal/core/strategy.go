package core

type Strategy interface {
	Info
	Strategy(*State) (Move, error)
}
