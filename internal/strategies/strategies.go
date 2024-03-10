package strategies

import "prisoners-dilemma/internal/core"

var Strategies []core.Strategy

func init() {
	Strategies = []core.Strategy{
		NewCuStrategy(),
		NewDuStrategy(),
		NewRANDStrategy(),
		NewTFTStrategy(),
		NewTFTTStrategy(),
		NewTTFTStrategy(),
		NewSTFTStrategy(),
		NewGrdTFTStrategy(),
		NewImpTFTStrategy(),
	}
}
