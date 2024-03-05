package game

import (
	"sync"
)

type Manager struct {
	runners []*Runner
}

func NewManager() *Manager {
	manager := Manager{
		runners: make([]*Runner, 0),
	}

	return &manager
}

func (m Manager) Runners() []*Runner {
	return m.runners
}

func (m Manager) Runner(i int) *Runner {
	return m.runners[i]
}

func (m *Manager) AddRunners(newRunners ...*Runner) {
	m.runners = append(m.runners, newRunners...)
}

func (m *Manager) RemoveRunner(i int) {
	m.runners = append(m.runners[:i], m.runners[i+1:]...)
}

func (m *Manager) Run(i int) (Report, error) {
	return m.runners[i].Run()
}

func (m *Manager) RunAll() ([]Report, error) {
	var wg sync.WaitGroup
	count := len(m.runners)
	reports := make([]Report, count)
	chans := make([]chan Report, count)

	for i := range m.runners {
		wg.Add(2)
		chans[i] = make(chan Report)

		go writer(m.runners[i], chans[i], &wg)
		go reader(&reports[i], chans[i], &wg)
	}

	wg.Wait()

	return reports, nil
}

func writer(runner *Runner, ch chan<- Report, wg *sync.WaitGroup) {
	defer close(ch)
	defer wg.Done()

	report, err := runner.Run()
	if err == nil {
		ch <- report
	}
}

func reader(report *Report, ch <-chan Report, wg *sync.WaitGroup) {
	defer wg.Done()

	for newReport := range ch {
		*report = newReport
	}
}
