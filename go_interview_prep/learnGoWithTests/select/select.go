package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func measureTimeResponse(url string) (time.Duration, error) {
	start := time.Now()
	res, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)

	return time.Since(start), nil
}

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		_, err := http.Get(url)
		if err != nil {
			return
		}
		close(ch)
	}()

	return ch
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

var tenSecondTimeout = 10 * time.Second
