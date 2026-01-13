package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Milua25/go_social/internal/store"
	"github.com/go-chi/chi/v5"
)

type postKey string

const postCTX postKey = "post"

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

// createPostHandler handles POST /v1/posts/create to create a post.
func (app *application) createPostHandler(w http.ResponseWriter, req *http.Request) {

	userID := 1

	var payload createPostPayload

	if err := readJSON(w, req, &payload); err != nil {
		app.badRequestError(w, req, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestError(w, req, err)
		return
	}

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

// getPostHandler returns a post along with its comments.
func (app *application) getPostHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	post_id, err := strconv.Atoi(chi.URLParam(req, "postID"))
	if err != nil {
		log.Println(err)
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

// deletePostHandler removes a post by id.
func (app *application) deletePostHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	post_id, err := strconv.Atoi(chi.URLParam(req, "postID"))
	if err != nil {
		log.Println(err)
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

// patchPostHandler updates mutable post fields using the context-loaded post.
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
			log.Println(err)
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
