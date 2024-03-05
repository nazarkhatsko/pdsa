package strategies

import "prisoners-dilemma/internal/core"

var Strategies []core.Strategy

func init() {
	Strategies = []core.Strategy{
		NewGoodStrategy(),
		NewRandStrategy(),
		NewTFTStrategy(),
	}
}
