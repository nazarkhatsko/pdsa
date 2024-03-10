package json_viewer

import (
	"encoding/json"
	"io"
	"prisoners-dilemma/internal/game"
)

func JSONViewer(reports []game.Report, w io.Writer) error {
	bytes, err := json.Marshal(reports)
	if err != nil {
		return err
	}
	w.Write(bytes)
	return nil
}
