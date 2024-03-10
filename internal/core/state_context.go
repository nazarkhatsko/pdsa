package core

type StateContext struct {
	balance Point
	moves   []Move
}

func (sc StateContext) Balance() Point {
	return sc.balance
}

func (sc StateContext) Moves() []Move {
	return sc.moves
}

func (sc StateContext) Move(i int) Move {
	if i > 0 {
		return sc.moves[i]
	}
	return sc.moves[len(sc.moves)-(i*-1)]
}

func (sc StateContext) FirstMove() (Move, bool) {
	l := len(sc.moves)
	if l > 0 {
		return sc.moves[0], true
	}
	return 0, false
}

func (sc StateContext) LastMove() (Move, bool) {
	l := len(sc.moves)
	if l > 0 {
		return sc.moves[l-1], true
	}
	return 0, false
}
