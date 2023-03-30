package ctx

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyResponseWriter struct {
	written bool
	body    string
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write(body []byte) (int, error) {
	s.written = true
	s.body = string(body)
	return 0, nil
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

type SpyStore struct {
	response string
	t        *testing.T
}

func mockWork(c context.Context, r string, data chan string) {
	var result string
	// simulate work by append each letter in `r` to the `result` string
	for _, cc := range r {
		select {
		case <-c.Done():
			fmt.Println("spy store cancelled")
			return
		default:
			time.Sleep(10 * time.Millisecond)
			result += string(cc)
		}
	}
	data <- result
}

func (s *SpyStore) Fetch(c context.Context) (string, error) {
	data := make(chan string, 1)

	go mockWork(c, s.response, data)

	// wait for mockWork() to finish or cancellation
	select {
	case <-c.Done():
		return "", c.Err()
	case res := <-data:
		return res, nil
	}
}

func TestServer(t *testing.T) {
	data := "hello world"

	t.Run(
		"returns data from store",
		func(t *testing.T) {
			store := newSpyStore(data, t)
			s := Server(store)

			request := httptest.NewRequest(http.MethodGet, "/", nil)
			cancellingCtx, _ := context.WithCancel(request.Context())
			request = request.WithContext(cancellingCtx)

			response := &SpyResponseWriter{}

			s.ServeHTTP(response, request)

			if !response.written {
				t.Error("expected a response to be written but it was not")
			}

			if response.body != data {
				t.Errorf("got %+q, want %+q", response.body, data)
			}

		})
	t.Run(
		"tells store to cancel work if request is cancelled",
		func(t *testing.T) {
			store := newSpyStore(data, t)
			s := Server(store)

			request := httptest.NewRequest(http.MethodGet, "/", nil)
			cancellingCtx, cancel := context.WithCancel(request.Context())
			time.AfterFunc(5*time.Millisecond, cancel)
			request = request.WithContext(cancellingCtx)

			response := &SpyResponseWriter{}

			s.ServeHTTP(response, request)

			if response.written {
				t.Errorf("a response was written but we cancelled the context")
			}
		},
	)
}

func newSpyStore(data string, t *testing.T) *SpyStore {
	return &SpyStore{data, t}
}

//func (s *SpyStore) assertWasCancelled() {
//	s.t.Helper()
//
//	if !s.cancelled {
//		s.t.Errorf("store was not told to cancel, but it should have been")
//	}
//}
//
//func (s *SpyStore) assertWasNotCancelled() {
//	s.t.Helper()
//
//	if s.cancelled {
//		s.t.Errorf("store was told to cancel, but it should not have been")
//	}
//}
