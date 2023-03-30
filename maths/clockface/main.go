package main

import (
	"github.com/chilledornaments/learn-go-with-tests/maths/svg"
	"os"
	"time"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
