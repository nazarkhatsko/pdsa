package game

import (
	"encoding/json"
	"prisoners-dilemma/internal/core"
	"strings"
	"time"
)

type Report struct {
	Tag            string           `json:"tag"`
	Iterations     int              `json:"iterations"`
	StartedAt      time.Time        `json:"startedAt"`
	FinishedAt     time.Time        `json:"finishedAt"`
	Solver         ReportInfo       `json:"solver"`
	WinnerStrategy string           `json:"winner_strategy"`
	Strategies     []ReportStrategy `json:"strategies"`
}

type ReportStrategy struct {
	ReportInfo
	Cooperates int          `json:"cooperates"`
	Defects    int          `json:"defects"`
	Balance    core.Point   `json:"balance"`
	Points     []core.Point `json:"points"`
	Moves      []core.Move  `json:"moves"`
}

type ReportInfo struct {
	ID          core.ID `json:"id"`
	Abbr        string  `json:"abbr"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
}

func (r Report) JSON(space int) string {
	bytes, err := json.MarshalIndent(r, "", strings.Repeat(" ", space))
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func ReportInfoFrom(info core.Info) ReportInfo {
	return ReportInfo{
		ID:          info.ID(),
		Abbr:        info.Abbr(),
		Name:        info.Name(),
		Description: info.Description(),
	}
}
