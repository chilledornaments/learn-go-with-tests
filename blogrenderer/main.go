package blogrenderer

import (
	"embed"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

// left off here https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/html-templates#back-to-the-real-work

type PostRenderer struct {
	templ    *template.Template
	mdParser *parser.Parser
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")

	if err != nil {
		return nil, err
	}
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	return &PostRenderer{templ: templ, mdParser: parser}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	return r.templ.ExecuteTemplate(w, "0_blog.gohtml", newPostVM(p, r))
}

type postViewModel struct {
	Post
	HTMLBody template.HTML
}

func newPostVM(p Post, r *PostRenderer) postViewModel {
	vm := postViewModel{Post: p}
	vm.HTMLBody = template.HTML(markdown.ToHTML([]byte(p.Body), r.mdParser, nil))
	return vm
}

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}
