package game

import (
	"prisoners-dilemma/internal/core"
	"prisoners-dilemma/internal/utils"
)

type Runner struct {
	id         core.ID
	tag        string
	iterations int
	solver     core.Solver
	strategies []core.Strategy
}

func NewRunner() *Runner {
	runner := Runner{
		id:         core.GenRandID(),
		tag:        "",
		iterations: 0,
		strategies: make([]core.Strategy, 0),
	}

	return &runner
}

func (r Runner) ID() core.ID {
	return r.id
}

func (r Runner) Tag() string {
	return r.tag
}

func (r Runner) Iterations() int {
	return r.iterations
}

func (r Runner) Solver() core.Solver {
	return r.solver
}

func (r Runner) Strategies() []core.Strategy {
	return r.strategies
}

func (r *Runner) SetTag(newTag string) {
	r.tag = newTag
}

func (r *Runner) SetIterations(newIterations int) {
	r.iterations = newIterations
}

func (r *Runner) SetSolver(newSolver core.Solver) {
	r.solver = newSolver
}

func (r *Runner) AddStrategies(newStrategies ...core.Strategy) {
	r.strategies = append(r.strategies, newStrategies...)
}

func (r *Runner) Run() (Report, error) {
	count := len(r.strategies)
	balances := make([]core.Point, count)
	points := make([][]core.Point, r.iterations)
	moves := make([][]core.Move, r.iterations)

	strategyIDs := utils.Map(r.strategies, func(s core.Strategy, _ int) core.ID {
		return s.ID()
	})

	finishedAt, startedAt, err := utils.WorkTime(func() error {
		stateManager := core.NewStateManager(strategyIDs)

		for i := 0; i < r.iterations; i++ {
			for _, strategy := range r.strategies {
				newMove, err := strategy.Strategy(stateManager.SwitchTo(strategy.ID()))
				if err != nil {
					return err
				}
				moves[i] = append(moves[i], newMove)
			}

			newPoints, err := r.solver.Solver(moves[i])
			if err != nil {
				return err
			}
			points[i] = newPoints
			for j, point := range points[i] {
				balances[j] += point
			}

			stateManager.Update(points[i], moves[i])
		}

		return nil
	})

	if err != nil {
		return Report{}, nil
	}

	report := Report{
		Tag:            r.tag,
		Iterations:     r.iterations,
		StartedAt:      startedAt,
		FinishedAt:     finishedAt,
		Solver:         ReportInfoFrom(r.solver),
		WinnerStrategy: "[NONE]",
		Strategies: utils.Map(r.strategies, func(s core.Strategy, i int) ReportStrategy {
			return ReportStrategy{
				ReportInfo: ReportInfoFrom(s),
				Cooperates: utils.Accumulate(moves, func(m []core.Move, acc int) int {
					if m[i] == core.MOVE_COOPERATE {
						return (acc + 1)
					}
					return acc
				}, 0),
				Defects: utils.Accumulate(moves, func(m []core.Move, acc int) int {
					if m[i] == core.MOVE_DEFECT {
						return (acc + 1)
					}
					return acc
				}, 0),
				Balance: balances[i],
				Points: utils.Accumulate(points, func(p []core.Point, acc []core.Point) []core.Point {
					return append(acc, p[i])
				}, make([]core.Point, 0)),
				Moves: utils.Accumulate(moves, func(m []core.Move, acc []core.Move) []core.Move {
					return append(acc, m[i])
				}, make([]core.Move, 0)),
			}
		}),
	}

	return report, nil
}
