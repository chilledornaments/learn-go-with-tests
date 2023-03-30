package di

import (
	"fmt"
	"io"
)

func Greet(w io.Writer, name string) {
	w.Write([]byte(fmt.Sprintf("Hello, %s", name)))
}
