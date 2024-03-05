package text_viewer

import (
	"io"
	"prisoners-dilemma/internal/game"
	"text/template"
)

func TextViewer(reports []game.Report, w io.Writer) error {
	funcs := template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
	}

	tmpl, err := template.New("layout.txt").Funcs(funcs).ParseFiles("internal/viewers/text/template/layout.txt")
	if err != nil {
		return err
	}

	err = tmpl.Execute(w, reports)
	if err != nil {
		return err
	}

	return nil
}
