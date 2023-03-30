package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/chilledornaments/blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, rustacean
---
A
B
C`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	want := blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World`,
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	got := posts[0]

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("expected number of posts to be %d but got %d", len(fs), len(posts))
	}

	assertPost(t, got, want)
}

//func TestFailingFS(t *testing.T) {
//	s := StubFailingFS{}
//
//	_, err := blogposts.NewPostsFromFS(s)
//
//	if err == nil {
//		t.Error("expected error but got none")
//	}
//}

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("error, always error")
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
