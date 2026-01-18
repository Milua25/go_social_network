package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Milua25/go_social/internal/store"
	"github.com/go-chi/chi/v5"
)

type postKey string

const postCTX postKey = "post"

// CreatePost godoc
//
//	@Summary		Creates a post
//	@Description	Creates a post
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		CreatePostPayload	true	"Post payload"
//	@Success		201		{object}	store.Post
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/posts [post]
//
// CreatePostPayload
type createPostPayload struct {
	Content string   `json:"content" validate:"required,max=1000"`
	Title   string   `json:"title" validate:"required,max=100"`
	Tags    []string `json:"tags"`
}

// type patchPayload struct {
// 	Title   string   `json:"title" validate:"required,max=100"`
// 	Content string   `json:"content" validate:"required,max=1000"`
// 	Tags    []string `json:"tags"`
// 	//Comments []string `json:"comments"`
// }

// createPostHandler godoc
//
//	@Summary		Create a post
//	@Description	Creates a post for the current user
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		createPostPayload	true	"Post payload"
//	@Success		200		{object}	store.Post
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/posts/create [post]
func (app *application) createPostHandler(w http.ResponseWriter, req *http.Request) {

	var payload createPostPayload

	if err := readJSON(w, req, &payload); err != nil {
		app.badRequestError(w, req, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestError(w, req, err)
		return
	}
	userID := getUserFromCtx(req).ID

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		UserID:  int64(userID),
		Tags:    payload.Tags,
	}

	ctx := req.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil {
		app.internalServerError(w, req, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, post); err != nil {
		app.internalServerError(w, req, err)
		return
	}
}

// getPostHandler godoc
//
//	@Summary		Get a post
//	@Description	Returns a post along with its comments
//	@Tags			posts
//	@Produce		json
//	@Param			postID	path		int	true	"Post ID"
//	@Success		200		{object}	store.Post
//	@Failure		404		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/posts/{postID} [get]
func (app *application) getPostHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	post_id, err := strconv.Atoi(chi.URLParam(req, "postID"))
	if err != nil {
		app.logger.Errorln(err)
		return
	}

	post, err := app.store.Posts.GetByID(ctx, post_id)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			app.notFoundError(w, req, err)
		default:
			app.internalServerError(w, req, err)
		}
		return
	}

	comments, err := app.store.Comments.GetByPostID(ctx, int64(post_id))
	if err != nil {
		app.internalServerError(w, req, err)
		return
	}

	post.Comments = make([]store.Comment, 0, len(comments))
	for _, c := range comments {
		if c == nil {
			continue
		}
		post.Comments = append(post.Comments, *c)
	}

	err = app.jsonResponse(w, http.StatusOK, post)
	if err != nil {
		app.internalServerError(w, req, err)
		return
	}

}

// deletePostHandler godoc
//
//	@Summary	Delete a post
//	@Tags		posts
//	@Param		postID	path	int	true	"Post ID"
//	@Success	204
//	@Failure	404	{object}	map[string]string
//	@Failure	500	{object}	map[string]string
//	@Router		/posts/{postID} [delete]
func (app *application) deletePostHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	post_id, err := strconv.Atoi(chi.URLParam(req, "postID"))
	if err != nil {
		app.logger.Errorln(err)
		return
	}

	err = app.store.Posts.DeleteByID(ctx, post_id)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			app.notFoundError(w, req, err)
		default:
			app.internalServerError(w, req, err)
		}
		return
	}
	app.jsonResponse(w, http.StatusNoContent, "Post deleted")
}

// patchPostHandler godoc
//
//	@Summary		Update a post
//	@Description	Updates post fields for the given ID
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			postID	path		int					true	"Post ID"
//	@Param			payload	body		createPostPayload	true	"Post payload"
//	@Success		200		{object}	store.Post
//	@Failure		400		{object}	map[string]string
//	@Failure		404		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/posts/{postID} [patch]
func (app *application) patchPostHandler(w http.ResponseWriter, req *http.Request) {

	var payload createPostPayload

	if err := readJSON(w, req, &payload); err != nil {
		app.badRequestError(w, req, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestError(w, req, err)
		return
	}

	post := getPostFromCtx(req)
	if post == nil {
		app.internalServerError(w, req, fmt.Errorf("post not in context"))
		return
	}

	// Update the post
	post.Title = payload.Title
	post.Content = payload.Content
	post.Tags = payload.Tags

	updatedPost, err := app.store.Posts.UpdateByID(req.Context(), post)
	if err != nil {
		app.internalServerError(w, req, err)
		return
	}
	err = app.jsonResponse(w, http.StatusOK, updatedPost)

	if err != nil {
		app.internalServerError(w, req, err)
		return
	}
}

// postsContextMiddleware fetches a post by path param and stores it in context.
func (app *application) postsContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		post_id, err := strconv.Atoi(chi.URLParam(req, "postID"))
		if err != nil {
			app.logger.Errorln(err)
			return
		}

		post, err := app.store.Posts.GetByID(ctx, post_id)

		if err != nil {
			switch err {
			case sql.ErrNoRows:
				app.notFoundError(w, req, err)
			default:
				app.internalServerError(w, req, err)
			}
			return
		}

		ctx = context.WithValue(ctx, postCTX, post)

		next.ServeHTTP(w, req.WithContext(ctx))
	})
}

// getPostFromCtx extracts the post placed in context by postsContextMiddleware.
func getPostFromCtx(req *http.Request) *store.Post {
	post, _ := req.Context().Value(postCTX).(*store.Post)

	return post
}
