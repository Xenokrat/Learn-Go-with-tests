package blogrenderer

import (
	"io"
	"html/template"
	"embed"
)

type Post struct {
	Title       string
	Body        string
	Description string
	Tags        []string
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func Render(w io.Writer, post Post) error {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err := templ.Execute(w, post); err != nil {
		return err
	}

	return nil
}
