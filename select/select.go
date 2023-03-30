package _select

import (
	"errors"
	"net/http"
	"time"
)

var defaultTimeout = 10 * time.Second

func Racer(a, b string) (w string, err error) {
	return ConfigurableRacer(a, b, defaultTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (w string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", errors.New("took longer than 10 seconds")
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{}, 1)
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
