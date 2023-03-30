package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func NewPostsFromFS(f fs.FS) ([]Post, error) {
	d, _ := fs.ReadDir(f, ".")

	var p []Post

	for _, file := range d {
		v, err := getPost(f, file.Name())

		if err != nil {
			return p, err
		}

		p = append(p, v)

	}

	return p, nil
}

func getPost(f fs.FS, fn string) (Post, error) {
	pf, err := f.Open(fn)
	if err != nil {
		return Post{}, err
	}

	defer pf.Close()

	return newPost(pf)
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagSeparator         = "Tags: "
)

func newPost(pf io.Reader) (Post, error) {
	scanner := bufio.NewScanner(pf)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	cleanTags := func(tags []string) []string {
		var v []string
		for _, t := range tags {
			v = append(v, strings.TrimSpace(t))
		}

		return v
	}

	// Move to next line
	title := readMetaLine(titleSeparator)
	description := readMetaLine(descriptionSeparator)
	tagString := readMetaLine(tagSeparator)
	// convert "foo, bar, whiz" to []string{}
	tags := cleanTags(strings.Split(tagString, ","))

	body := readBody(scanner)

	return Post{Title: title, Description: description, Tags: tags, Body: body}, nil
}

func readBody(s *bufio.Scanner) string {
	// Skip "---" line
	s.Scan()

	bodyBuffer := bytes.Buffer{}
	// scan until there's nothing left
	for s.Scan() {
		fmt.Fprintln(&bodyBuffer, s.Text())
	}
	return strings.TrimSuffix(bodyBuffer.String(), "\n")

}
