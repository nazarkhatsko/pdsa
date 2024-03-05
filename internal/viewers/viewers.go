package viewers

import (
	"io"
	"prisoners-dilemma/internal/game"
	html_viewer "prisoners-dilemma/internal/viewers/html"
	json_viewer "prisoners-dilemma/internal/viewers/json"
	text_viewer "prisoners-dilemma/internal/viewers/text"
)

type Viewer func([]game.Report, io.Writer) error

var Viewers map[string]Viewer

func init() {
	Viewers = map[string]Viewer{
		"html": html_viewer.HTMLViewer,
		"json": json_viewer.JSONViewer,
		"text": text_viewer.TextViewer,
	}
}
