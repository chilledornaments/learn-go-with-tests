package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("verify correct url is returned", func(t *testing.T) {
		slowServer := makeDelayedServer(10 * time.Millisecond)
		fastServer := makeDelayedServer(5 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		want := fastServer.URL

		got, _ := Racer(slowServer.URL, fastServer.URL)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

	t.Run("verify error after 10 seconds", func(t *testing.T) {
		s := makeDelayedServer(3 * time.Second)

		defer s.Close()

		_, err := ConfigurableRacer(s.URL, s.URL, 1*time.Second)

		if err == nil {
			t.Errorf("did not receive error but expected one")
		}
	})

}

func makeDelayedServer(d time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(d)
		writer.WriteHeader(http.StatusOK)
	}))
}
