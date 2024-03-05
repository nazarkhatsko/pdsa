package core

import "math/rand"

type ID uint64

func GenRandID() ID {
	return ID(rand.Uint64())
}
