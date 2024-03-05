package html_viewer

import (
	"html/template"
	"io"
	"prisoners-dilemma/internal/game"
	"time"
)

func HTMLViewer(reports []game.Report, w io.Writer) error {
	funcs := template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
		"iterate": func(count int) []int {
			var i int
			var items []int
			for i = 0; i < count; i++ {
				items = append(items, i)
			}
			return items
		},
		"date": func(t time.Time) string {
			return t.Format(time.ANSIC)
		},
	}

	tmpl, err := template.New("layout.html").Funcs(funcs).ParseFiles("internal/viewers/html/template/layout.html")
	if err != nil {
		return err
	}

	tmpl.New("jquery.min.js").ParseFiles("internal/viewers/html/template/packages/jquery/jquery.min.js")

	tmpl.New("bootstrap.min.css").ParseFiles("internal/viewers/html/template/packages/bootstrap/bootstrap.min.css")
	tmpl.New("bootstrap.bundle.min.js").ParseFiles("internal/viewers/html/template/packages/bootstrap/bootstrap.bundle.min.js")

	tmpl.New("theme.js").ParseFiles("internal/viewers/html/template/scripts/theme.js")
	tmpl.New("search.js").ParseFiles("internal/viewers/html/template/scripts/search.js")
	tmpl.New("popover.js").ParseFiles("internal/viewers/html/template/scripts/popover.js")

	err = tmpl.Execute(w, reports)
	if err != nil {
		return err
	}

	return nil
}
