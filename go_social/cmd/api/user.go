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

// getUserByIdHandler godoc
//
//	@Summary		Get user
//	@Description	Returns a user by ID
//	@Tags			users
//	@Produce		json
//	@Param			userID	path		int	true	"User ID"
//	@Success		200		{object}	store.User
//	@Failure		404		{object}	map[string]string
//	@Router			/users/{userID} [get]
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

// getAllUsersHandler godoc
//
//	@Summary		List users
//	@Description	Return all users
//	@Tags			users
//	@Produce		json
//	@Success		200	{array}	store.User
//	@Router			/users [get]
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

// followUserHandler godoc
//
//	@Summary		Follow user
//	@Description	Current user follows the target user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			userID		path	int			true	"User ID"
//	@Param			follower	body	Follower	true	"Follower payload"
//	@Success		204
//	@Failure		409	{object}	map[string]string
//	@Router			/users/{userID}/follow [put]
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

// unfollowUserHandler godoc
//
//	@Summary		Unfollow user
//	@Description	Current user unfollows the target user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			userID		path	int			true	"User ID"
//	@Param			follower	body	Follower	true	"Unfollow payload"
//	@Success		204
//	@Router			/users/{userID}/unfollow [put]
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

// activateUserHandler godoc
//
//	@Summary		Activate user
//	@Description	Activates a user using an invitation token
//	@Tags			users
//	@Produce		json
//	@Param			token	path	string	true	"Invitation token"
//	@Success		204
//	@Failure		400	{object}	map[string]string
//	@Router			/users/activate/{token} [post]
func (app *application) activateUserHandler(w http.ResponseWriter, req *http.Request) {
	token := chi.URLParam(req, "token")

	err := app.store.Users.Activate(req.Context(), token)

	if err != nil {
		switch err {
		case store.ErrNotFound:
			app.badRequestError(w, req, err)
		default:
			app.internalServerError(w, req, err)
		}
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, req, err)
	}

}
