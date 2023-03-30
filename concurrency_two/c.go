package concurrency_two

import (
	"fmt"
	"math/rand"
	"time"
)

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()")

func work(c chan rune) {
	r := rand.Intn(len(chars))

	c <- chars[r]
	return
}

func generator() {
	rand.Seed(time.Now().UnixNano())
	var s string
	ch := make(chan rune)

	for i := 0; i < 100; i++ {
		go work(ch)
		//x := <-ch
		//s += string(x)
	}

	for i := 0; i < 100; i++ {
		x := <-ch
		s += string(x)
	}

	fmt.Println(s)
}
