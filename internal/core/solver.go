package core

type Solver interface {
	Info
	Solver([]Move) ([]Point, error)
}
