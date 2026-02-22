package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("not using mock test", func(t *testing.T) {
		slowURL := "https://www.facebook.com"
		fastURL := "https://www.quii.dev"

		want := fastURL
		got, _ := Racer(slowURL, fastURL)
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("using mock test", func(t *testing.T) {
		slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			time.Sleep(20 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
		}))

		fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
		fastURL := fastServer.URL
		slowURL := slowServer.URL

		want := fastURL

		got, _ := Racer(slowURL, fastURL)
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		slowServer.Close()
		fastServer.Close()
	})

	//t.Run("using mock test2", func(t *testing.T) {
	//	slowServer := makeDelayedServer(20 * time.Millisecond)
	//	fastServer := makeDelayedServer(0 * time.Millisecond)
	//	defer slowServer.Close()
	//	defer fastServer.Close()
	//
	//	slowURL := slowServer.URL
	//	fastURL := fastServer.URL
	//
	//	want := fastURL
	//	got := Racer(slowURL, fastURL)
	//
	//	if got != want {
	//		t.Errorf("got %q, want %q", got, want)
	//	}
	//})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Second)
		serverB := makeDelayedServer(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()
		_, err := Racer(serverA.URL, serverB.URL)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})

	t.Run("returns an error if a server doesn't respond within the specified time", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})

}

func makeDelayedServer(duration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))

}
