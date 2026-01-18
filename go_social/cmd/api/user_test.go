package main

import (
	"net/http"
	"testing"

	"github.com/Milua25/go_social/internal/store/cache"
	"github.com/stretchr/testify/mock"
)

func TestGetAllUsersHandler(t *testing.T) {
	app := newTestApplication(t)
	mux := app.mount()
	testToken, err := app.authenticatior.GenerateToken(nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("should not allow unauthenticated requests", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/v1/users/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := execRequest(req, mux)

		checkResponseCode(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("should allow unauthenticated requests", func(t *testing.T) {
		mockCacheStore := app.cacheStorage.Users.(*cache.MockeUserStore)

		mockCacheStore.On("Get", int64(1)).Return(nil, nil).Twice()
		mockCacheStore.On("Get", mock.Anything).Return(nil)

		req, err := http.NewRequest(http.MethodGet, "/v1/users/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Authorization", "Bearer "+testToken)

		rr := execRequest(req, mux)

		checkResponseCode(t, http.StatusOK, rr.Code)
		mockCacheStore.Calls = nil // Reset mock expectations
	})

	t.Run("should hit the cache first and if not exists it sets the user on the cache", func(t *testing.T) {
		mockCacheStore := app.cacheStorage.Users.(*cache.MockeUserStore)

		mockCacheStore.On("Get", int64(42)).Return(nil, nil)
		mockCacheStore.On("Get", int64(1)).Return(nil, nil)
		mockCacheStore.On("Set", mock.Anything, mock.Anything).Return(nil)

		req, err := http.NewRequest(http.MethodGet, "/v1/users/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Authorization", "Bearer "+testToken)

		rr := execRequest(req, mux)

		checkResponseCode(t, http.StatusOK, rr.Code)

		mockCacheStore.AssertNumberOfCalls(t, "Get", 2)

		mockCacheStore.Calls = nil // Reset mock expectations
	})
}
