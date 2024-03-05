package app

import (
	"fmt"
	"os"
	"prisoners-dilemma/internal/config"
	"prisoners-dilemma/internal/core"
	"prisoners-dilemma/internal/game"
	"prisoners-dilemma/internal/solvers"
	"prisoners-dilemma/internal/strategies"
	"prisoners-dilemma/internal/viewers"
	"slices"
	"strings"
)

type App struct {
	manager *game.Manager
	viewers []func([]game.Report) error
}

func NewApp(conf *config.Config) *App {
	app := App{
		manager: game.NewManager(),
		viewers: make([]func([]game.Report) error, 0),
	}

	for _, r := range conf.Runners {
		runner := game.NewRunner()
		runner.SetTag(r.Tag)
		runner.SetIterations(r.Iterations)

		index := slices.IndexFunc(solvers.Solvers, func(s core.Solver) bool {
			return strings.EqualFold(s.Abbr(), r.Solver)
		})
		if index == -1 {
			panic(fmt.Sprintf("error: solver with abbr \"%s\" is not found", r.Solver))
		}
		runner.SetSolver(solvers.Solvers[index])

		for _, strategy := range r.Strategies {
			index := slices.IndexFunc(strategies.Strategies, func(s core.Strategy) bool {
				return strings.EqualFold(s.Abbr(), strategy)
			})
			if index == -1 {
				panic(fmt.Sprintf("error: strategy with abbr \"%s\" is not found", strategy))
			}
			runner.AddStrategies(strategies.Strategies[index])
		}

		app.manager.AddRunners(runner)
	}

	for _, v := range conf.Viewers {
		viewer, ok := viewers.Viewers[v.Type]
		if !ok {
			panic("error: inccorect viewer type")
		}

		out, err := os.Create(v.Out)
		if err != nil {
			panic(err)
		}

		app.viewers = append(app.viewers, func(reports []game.Report) error {
			defer out.Close()

			err = viewer(reports, out)
			if err != nil {
				return err
			}

			return nil
		})
	}

	return &app
}

func (a *App) Run() error {
	reports, err := a.manager.RunAll()
	if err != nil {
		return err
	}

	for _, viewer := range a.viewers {
		err := viewer(reports)
		if err != nil {
			return err
		}
	}

	return nil
}
