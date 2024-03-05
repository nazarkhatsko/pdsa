package json_viewer

import (
	"encoding/json"
	"io"
	"prisoners-dilemma/internal/game"
	"strings"
)

func JSONViewer(reports []game.Report, w io.Writer) error {
	bytes, err := json.MarshalIndent(reports, "", strings.Repeat(" ", 4))
	if err != nil {
		return err
	}
	w.Write(bytes)
	return nil
}
