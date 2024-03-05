package core

type Move uint32

const (
	MOVE_NULL      Move = 0
	MOVE_COOPERATE Move = 1
	MOVE_DEFECT    Move = 2
)
