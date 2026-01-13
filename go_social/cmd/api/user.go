package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Milua25/go_social/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/lib/pq"
)

type userKey string

const userKeyCtx userKey = "userID"

type Follower struct {
	FollowerID int64 `json:"follower_id" validate:"required,min=1"`
}

// getUserByIdHandler handles GET /users/{userID} and returns a single user.
func (app *application) getUserByIdHandler(w http.ResponseWriter, req *http.Request) {

	user_id_string := chi.URLParam(req, "userID")
	if user_id_string == "" {
		app.badRequestError(w, req, fmt.Errorf("missing user id"))
		return
	}

	user_id, err := strconv.Atoi(user_id_string)
	if err != nil {
		app.badRequestError(w, req, err)
		return
	}

	ctx := req.Context()

	user, err := app.store.Users.GetUserByID(ctx, user_id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			app.notFoundError(w, req, fmt.Errorf("No user with that id"))
			return
		default:
			app.badRequestError(w, req, err)
			return
		}
	}
	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, req, err)
		return
	}
}

// getAllUsersHandler handles GET /users and returns all users.
func (app *application) getAllUsersHandler(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()

	users, err := app.store.Users.GetUsers(ctx)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			app.notFoundError(w, req, fmt.Errorf("No users"))
			return
		default:
			app.badRequestError(w, req, err)
			return
		}
	}
	if err := app.jsonResponse(w, http.StatusOK, users); err != nil {
		app.internalServerError(w, req, err)
		return
	}
}

// followUserHandler handles following a user from the current requester.
func (app *application) followUserHandler(w http.ResponseWriter, req *http.Request) {
	// get id
	user := getUserFromCtx(req)

	// revert to use auth
	var payload Follower

	if err := readJSON(w, req, &payload); err != nil {
		app.badRequestError(w, req, err)
		return
	}

	err := Validate.Struct(payload)
	if err != nil {
		app.badRequestError(w, req, err)
		return
	}

	ctx := req.Context()

	// context
	if err := app.store.Followers.Follow(ctx, payload.FollowerID, user.ID); err != nil {
		var pqErr *pq.Error

		switch {
		case errors.As(err, &pqErr) && pqErr.Code == "23505" && pqErr.Constraint == "followers_pkey":
			app.conflictResponseError(w, req, fmt.Errorf("already following this user"))
			return
		default:
			app.internalServerError(w, req, err)
			return
		}
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, req, err)
		return
	}
}

// unfollowUserHandler handles removing a follow from the current requester.
func (app *application) unfollowUserHandler(w http.ResponseWriter, req *http.Request) {

	// get id
	user := getUserFromCtx(req)

	// revert to use auth
	var payload Follower

	if err := readJSON(w, req, &payload); err != nil {
		app.badRequestError(w, req, err)
		return
	}
	err := Validate.Struct(payload)
	if err != nil {
		app.badRequestError(w, req, err)
		return
	}

	ctx := req.Context()

	// context
	if err := app.store.Followers.UnFollow(ctx, payload.FollowerID, user.ID); err != nil {
		app.internalServerError(w, req, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, req, err)
		return
	}
}

// Setting up middleware for the getUserID
// getUserContextIdMiddleware loads the route user into context for downstream handlers.
func (app *application) getUserContextIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user_id_string := chi.URLParam(r, "userID")
		if user_id_string == "" {
			app.badRequestError(w, r, fmt.Errorf("missing user id"))
			return
		}

		user_id, err := strconv.Atoi(user_id_string)
		if err != nil {
			app.badRequestError(w, r, err)
			return
		}
		ctx := r.Context()

		user, err := app.store.Users.GetUserByID(ctx, user_id)
		if err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				app.notFoundError(w, r, fmt.Errorf("No user with that id"))
				return
			default:
				app.badRequestError(w, r, err)
				return
			}
		}

		ctx = context.WithValue(r.Context(), userKeyCtx, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// getUserFromCtx extracts the user stored in context by getUserContextIdMiddleware.
func getUserFromCtx(req *http.Request) *store.User {
	val, ok := req.Context().Value(userKeyCtx).(*store.User)
	if !ok {
		//missing
		log.Println("user id key missing")
		return nil
	}

	return val
}
