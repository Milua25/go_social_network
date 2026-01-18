package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Milua25/go_social/internal/auth"
	"github.com/Milua25/go_social/internal/ratelimiter"
	"github.com/Milua25/go_social/internal/store"
	"github.com/Milua25/go_social/internal/store/cache"
	"go.uber.org/zap"
)

func newTestApplication(t *testing.T, cfg config) *application {
	t.Helper()

	logger := zap.NewNop().Sugar()
	// Uncomment to enable logs
	// logger := zap.Must(zap.NewProduction()).Sugar()
	mockStore := store.NewMockStore()
	mockCacheStore := cache.NewMockStore()

	testAuth := &auth.TestAuthenticator{}

	// Rate limiter
	rateLimiter := ratelimiter.NewFixedWindowLimiter(
		cfg.rateLimiter.RequestsPerTimeFrame,
		cfg.rateLimiter.TimeFrame,
	)

	return &application{
		logger:         logger,
		store:          mockStore,
		cacheStorage:   mockCacheStore,
		authenticatior: testAuth,
		config:         cfg,
		rateLimiter:    rateLimiter,
	}
}

func execRequest(req *http.Request, mux http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("expected response code to be %d and we got %d", expected, actual)
	}
}
