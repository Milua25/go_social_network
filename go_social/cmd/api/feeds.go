package main

import (
	"net/http"

	"github.com/Milua25/go_social/internal/store"
)

func (app *application) getUserFeedHandler(w http.ResponseWriter, req *http.Request) {
	// pagination
	fq := store.PaginatedFeedQuery{
		Limit:  20,
		Offset: 0,
		Sort:   "desc",
	}

	fq, err := fq.Parse(req)
	if err != nil {
		app.badRequestError(w, req, err)
		return
	}

	if err := Validate.Struct(fq); err != nil {
		app.badRequestError(w, req, err)
		return
	}

	ctx := req.Context()

	posts, err := app.store.Posts.GetUserFeed(ctx, int64(41), fq)

	if err != nil {
		app.internalServerError(w, req, err)
		return
	}

	app.jsonResponse(w, http.StatusOK, posts)

	// filters

	// search
}
