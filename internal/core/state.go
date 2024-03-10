package core

type State struct {
	iteration int
	self      StateSelf
	oponents  []StateOponent
}

func NewState() *State {
	state := State{
		iteration: 0,
		self:      StateSelf{},
		oponents:  make([]StateOponent, 0),
	}

	return &state
}

func (s State) Iteration() int {
	return s.iteration
}

func (s State) Self() *StateSelf {
	return &s.self
}

func (s State) OponentsCount() int {
	return len(s.oponents)
}

func (s State) Oponents() []StateOponent {
	return s.oponents
}

func (s State) FirstOponent() StateOponent {
	return s.oponents[0]
}

func (s State) LastOponent() StateOponent {
	return s.oponents[len(s.oponents)-1]
}
