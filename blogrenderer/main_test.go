package blogrenderer_test

import (
	"bytes"
	approvals "github.com/approvals/go-approval-tests"
	"github.com/chilledornaments/learn-go-with-tests/blogrenderer"
	"io"
	"testing"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := blogrenderer.NewPostRenderer()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("post converted to HTML", func(t *testing.T) {
		b := bytes.Buffer{}
		err := postRenderer.Render(&b, aPost)

		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, b.String())
	})

}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := blogrenderer.NewPostRenderer()

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}
