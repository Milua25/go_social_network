package main

import (
	"net/http"

	"github.com/Milua25/go_social/internal/store"
)

// getUserFeedHandler godoc
//
//	@Summary		Fetches the user feed
//	@Description	Fetches the user feed
//	@Tags			feed
//	@Accept			json
//	@Produce		json
//	@Param			since	query		string	false	"Since"
//	@Param			until	query		string	false	"Until"
//	@Param			limit	query		int		false	"Limit"
//	@Param			offset	query		int		false	"Offset"
//	@Param			sort	query		string	false	"Sort"
//	@Param			tags	query		string	false	"Tags"
//	@Param			search	query		string	false	"Search"
//	@Success		200		{object}	[]store.PostWithMetadata
//	@Failure		400		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/users/feed [get]
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
